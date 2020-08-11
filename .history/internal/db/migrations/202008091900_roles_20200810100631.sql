\set ON_ERROR_STOP on

BEGIN;
    DO $$
    DECLARE migrationId VARCHAR := '202008091900_roles';   
    DECLARE usr text := '';
    DECLARE admin_usr text := '';
    BEGIN  
        IF NOT (SELECT curator.check_migration(migrationId)) THEN
          ------------------------------------------------------------
            --migration script start
            EXECUTE  'EXECUTE get_v1 ' INTO usr;            
            EXECUTE  'EXECUTE get_v2 ' INTO admin_usr;
            ------------------------------------------------------------
            -- SCHEMA:
            BEGIN
                CREATE TABLE deputy.roles (
                    id serial PRIMARY KEY,
                    name CHARACTER VARYING NOT NULL,
                    parent_id BIGINT
                );
                
                CREATE INDEX ON roles(id);
                
                ALTER TABLE roles
                    ADD CONSTRAINT role_parent_fk FOREIGN KEY parent_id REFERENCES roles(id);
                
            END;
            EXECUTE 'ALTER TABLE deputy.roles OWNER TO ' || quote_ident(admin_usr);
            EXECUTE 'GRANT ALL ON TABLE deputy.roles TO ' || quote_ident(usr); 

            CREATE INDEX stat_events_start_time_idx ON curator.stat_events USING btree (org_id, space_id, workspace_id, experience_id, scenario_id, start_time DESC);
            CREATE INDEX stat_events_served_idx ON curator.stat_events USING btree (org_id, space_id, workspace_id, experience_id, scenario_id, start_time DESC) WHERE ((action)::text = 'served'::text AND category='sdk');
            ---------------------------------------------------------------------------------------------------
            -- This function returuns all required buckets for the requested period. Each row represents one bucket for one metric (action/category)
            -- plus one NULL bucket for each metric (action/category) which represents the previous period total 
            CREATE OR REPLACE FUNCTION curator.get_scenario_metrics_buckets(
                orgId VARCHAR(40), 
                spaceId VARCHAR(40), 
                workspaceId VARCHAR(40),
                experienceId VARCHAR(40), 
                scenarioId VARCHAR(40), 
                categorynames VARCHAR[], 
                actionnames VARCHAR[], 
                
                -- we no operate in whole minutes, not seconds:
                timeperiodmins INT, 
                timebucketmins INT

                )
             RETURNS TABLE(
                bucket TIMESTAMP, 
                category VARCHAR,
                action VARCHAR, 
                activity_count BIGINT)
                --is_previous BOOLEAN)
             LANGUAGE plpgsql
            AS $function$
                DECLARE
                    -- now truncated to a round minute:
                     last_activity_time TIMESTAMP := date_trunc('minute', NOW() AT TIME ZONE 'UTC');
                     first_activity_time TIMESTAMP := last_activity_time - (INTERVAL '1 minute' * timeperiodmins);
                     
                     -- start time of the previous period:
                     prev_first_activity_time TIMESTAMP := first_activity_time - (INTERVAL '1 minute' * timeperiodmins);
                     
                     
                     last_activity_time_epoch BIGINT := extract('epoch' from last_activity_time);
                     first_activity_time_epoch BIGINT := extract('epoch' from first_activity_time);
                     
                     timebucketsecs INT := timebucketmins * 60;
                -- 24h:  + 15mins - 900 seconds 
                BEGIN
                -- make sure timebucketmins > 0
                IF timebucketmins <= 0 THEN
                    RAISE invalid_parameter_value USING MESSAGE = 'timebucketmins must be a postive number';
                    RETURN;
                END IF;
               
                -- make sure timeperiodmins > 0
                IF timeperiodmins <= 0 THEN
                    RAISE invalid_parameter_value USING MESSAGE = 'timeperiodmins must be a postive number';
                    RETURN;
                END IF;
               
                IF timebucketmins > timeperiodmins THEN
                    RAISE invalid_parameter_value USING MESSAGE = 'timebucketmins must be a greater than timeperiodmins';
                    RETURN;
                END IF;
                
                IF array_length(categorynames, 1) <> array_length(actionnames, 1) THEN
                    RAISE invalid_parameter_value USING MESSAGE = 'categorynames and actionnames must have the same length';
                    RETURN;
                END IF;
               
                    
                RETURN QUERY  
                WITH buckets AS (
                    SELECT bucket_time
                    FROM generate_series(first_activity_time_epoch, last_activity_time_epoch - 1, timebucketsecs) gs(bucket_time) 
                ),
                non_empty_activity AS (
                    SELECT
                        first_activity_time_epoch + floor(((extract('epoch' from e.start_time) - first_activity_time_epoch) / timebucketsecs )) * timebucketsecs AS bucket_time,
                        mids.category,
                        mids."action",
                        sum(event_count)::BIGINT  AS sum_activity_count
                    FROM unnest(categoryNames, actionNames) mids (category, action)
                    LEFT JOIN  curator.stat_events e
                        ON (
                                (mids.category = e.category OR (mids.category IS NULL AND e.category IS NULL))
                                AND mids."action" = e."action"
                            )
                    WHERE 
                        e.org_Id = orgId
                        AND e.space_id = spaceId
                        AND workspace_id = workspaceId
                        AND	e.experience_id = experienceId
                        AND e.scenario_id = scenarioId

                        AND e.start_time >= first_activity_time 
                        AND e.end_time <= last_activity_time

                    GROUP BY mids.category, mids."action", bucket_time
                ),
                non_empty_prev_counts AS (
                    SELECT
                        mids.category,
                        mids."action",
                        sum(event_count)::BIGINT  AS prev_sum_activity_count
                    FROM unnest(categoryNames, actionNames) mids (category, action)
                    LEFT JOIN  curator.stat_events e
                        ON (
                                (mids.category = e.category OR (mids.category IS NULL AND e.category IS NULL))
                                AND mids."action" = e."action"
                            )
                    WHERE 
                        e.org_Id = orgId
                        AND e.space_id = spaceId
                        AND workspace_id = workspaceId
                        AND	e.experience_id = experienceId
                        AND e.scenario_id = scenarioId

                        AND e.start_time >= prev_first_activity_time 
                        AND e.end_time <= first_activity_time

                    GROUP BY mids.category, mids."action"
                )
                    SELECT TO_TIMESTAMP(b.bucket_time) AT TIME ZONE 'UTC' AS bucket,
                        mids.category,
                        mids."action",
                        COALESCE(nec.sum_activity_count, 0) AS activity_count
                    FROM UNNEST(categoryNames, actionNames) mids (category, action)
                    CROSS JOIN buckets b  -- making sure ALL buckets appear in the results for EVERY action+category even if they are zeros
                    LEFT JOIN non_empty_activity nec ON (
                                                            (mids.category = nec.category OR (mids.category IS NULL AND nec.category IS NULL))
                                                            AND mids."action" = nec."action"
                                                            AND nec.bucket_time = b.bucket_time
                                                    )

                -- prev period counts will be returned as null timestamps:
                UNION
                    SELECT NULL AS bucket,
                        mids.category,
                        mids."action",
                        COALESCE(pc.prev_sum_activity_count, 0) AS activity_count
                    FROM UNNEST(categoryNames, actionNames) mids (category, action)
                        LEFT JOIN non_empty_prev_counts pc ON (
                                                            (mids.category = pc.category OR (mids.category IS NULL AND pc.category IS NULL))
                                                            AND mids."action" = pc."action"
                                                            )
                
                ORDER BY bucket ASC;


                END;
            $function$;
            
            EXECUTE 'ALTER FUNCTION curator.get_scenario_metrics_buckets(VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR[], VARCHAR[], INT, INT)  OWNER TO ' || quote_ident(admin_usr);
            EXECUTE 'GRANT ALL ON FUNCTION curator.get_scenario_metrics_buckets(VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR[], VARCHAR[], INT, INT) TO ' || quote_ident(usr);
            ---------------------------------------------------------------------------------------------------
            -- This function reformats the data returned by curator.get_scenario_metrics_buckets so that each row represents ALL the data for each metric (action/category)
            -- the returned activity_counts holds all the bucket totals for the requested period\
            -- curr_total represents the total count for the entire current period (e.g current 30 days)
            -- prev_total represents the total count for the entire previous period (e.g previous 30 days)
            CREATE OR REPLACE FUNCTION curator.get_scenario_metrics(
                orgId VARCHAR(40), 
                spaceId VARCHAR(40), 
                workspaceId VARCHAR(40),
                experienceId VARCHAR(40), 
                scenarioId VARCHAR(40), 
                categorynames VARCHAR[], 
                actionnames VARCHAR[], 
                
                -- we no operate in whole minutes, not seconds:
                timeperiodmins INT, 
                timebucketmins INT

                )
             RETURNS TABLE(
                category VARCHAR,
                "action" VARCHAR, 
                start_time TIMESTAMP, 
                activity_counts BIGINT[],
                curr_total BIGINT,
                prev_total BIGINT)
             LANGUAGE plpgsql
            AS $function$
                BEGIN
                
                    RETURN QUERY  
                    SELECT 
                        gsm.category, 
                        gsm."action",
                        MIN(gsm.bucket) AS start_time, 
                        ARRAY_AGG(gsm.activity_count ORDER BY bucket ASC) FILTER (WHERE bucket IS NOT NULL) AS activity_counts, 
                        
                        -- Only count the rows WITHOUT nulls in them (=current period buckets)
                        (SUM(gsm.activity_count) FILTER (WHERE bucket IS NOT NULL))::BIGINT  curr_total,
                        
                        -- Only count the rows with nulls in them (=previous period totals)
                        (SUM(gsm.activity_count) FILTER (WHERE bucket IS NULL))::BIGINT  AS prev_total
                    FROM curator.get_scenario_metrics_buckets(orgId, spaceId, workspaceId, experienceId, scenarioId, categorynames, actionnames, timeperiodmins, timebucketmins) gsm
                    GROUP BY gsm.category, gsm."action";
                END;
            $function$;




            EXECUTE 'ALTER FUNCTION curator.get_scenario_metrics(VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR[], VARCHAR[], INT, INT)  OWNER TO ' || quote_ident(admin_usr);
            EXECUTE 'GRANT ALL ON FUNCTION curator.get_scenario_metrics(VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR(40), VARCHAR[], VARCHAR[], INT, INT) TO ' || quote_ident(usr);
            ---------------------------------------------------------------------------------------------------


            ------------------------------------------------------------  
            --migration script end
            PERFORM curator.create_migration(migrationId);
            ------------------------------------------------------------
        END IF;
    END    
    $$;
COMMIT;


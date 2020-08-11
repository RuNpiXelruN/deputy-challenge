\set ON_ERROR_STOP on

BEGIN;
    DO $$
    DECLARE migrationId VARCHAR := '202008092000_get_subordinates';   
    DECLARE usr text := '';
    DECLARE admin_usr text := '';
    BEGIN  
        IF NOT (SELECT check_migration(migrationId)) THEN
          ------------------------------------------------------------
            --migration script start
            EXECUTE  'EXECUTE get_dbUser ' INTO usr;            
            EXECUTE  'EXECUTE get_dbAdmin ' INTO admin_usr;
            ------------------------------------------------------------
            -- This function returuns all required buckets for the requested period. Each row represents one bucket for one metric (action/category)
            -- plus one NULL bucket for each metric (action/category) which represents the previous period total
            DROP FUNCTION IF EXISTS getSubordinates(userId integer);
            
            CREATE OR REPLACE FUNCTION getSubordinates(userid integer)
            RETURNS SETOF users
            LANGUAGE plpgsql
            AS 
            $func$
                BEGIN		
                    RETURN QUERY (			
                        WITH RECURSIVE
                            user_role_id AS (
                                SELECT u.role_id
                                FROM users AS u
                                WHERE u.id = userId
                            ),
                            sub_tree AS (
                                SELECT id, name, 1 AS relative_depth
                                FROM roles
                                WHERE roles.id = (SELECT * FROM user_role_id)
                                UNION ALL
                                SELECT r.id, r."name", st.relative_depth + 1
                                FROM roles AS r, sub_tree AS st
                                WHERE r.parent_id = st.id
                            )
                        SELECT u.id, u."name", u.role_id AS role
                        FROM users u, sub_tree st
                        WHERE u.role_id = st.id
                        AND st.relative_depth > 1
                    );
                END;
            $func$;            
            ------------------------------------------------------------  
            --migration script end
            PERFORM create_migration(migrationId);
            ------------------------------------------------------------
        END IF;
    END    
    $$;
COMMIT;
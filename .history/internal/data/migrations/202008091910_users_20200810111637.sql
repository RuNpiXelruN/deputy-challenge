\set ON_ERROR_STOP on

BEGIN;
    DO $$
    DECLARE migrationId VARCHAR := '202008091910_users';   
    DECLARE usr text := '';
    DECLARE admin_usr text := '';
    BEGIN  
        IF NOT (SELECT deputy.check_migration(migrationId)) THEN
          ------------------------------------------------------------
            --migration script start
            EXECUTE  'EXECUTE get_dbUser ' INTO usr;            
            EXECUTE  'EXECUTE get_dbAdmin ' INTO admin_usr;
            ------------------------------------------------------------
            -- SCHEMA:
            BEGIN
                CREATE TABLE deputy.users (
                    id serial PRIMARY KEY,
                    name CHARACTER VARYING NOT NULL,
                    role_id BIGINT NOT NULL
                );                            
                
            END;
            EXECUTE 'ALTER TABLE deputy.roles OWNER TO ' || quote_ident(admin_usr);
            EXECUTE 'GRANT ALL ON TABLE deputy.roles TO ' || quote_ident(usr); 

            CREATE INDEX ON deputy.users(role_id);
            ALTER TABLE deputy.users
                ADD CONSTRAINT user_role_fk FOREIGN KEY (role_id) REFERENCES deputy.roles(id);

            ---------------------------------------------------------------------------------------------------
            -- This function returuns all required buckets for the requested period. Each row represents one bucket for one metric (action/category)
            -- plus one NULL bucket for each metric (action/category) which represents the previous period total 
            
            ------------------------------------------------------------  
            --migration script end
            PERFORM deputy.create_migration(migrationId);
            ------------------------------------------------------------
        END IF;
    END    
    $$;
COMMIT;
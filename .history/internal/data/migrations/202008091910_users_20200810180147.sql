\set ON_ERROR_STOP on

BEGIN;
    DO $$
    DECLARE migrationId VARCHAR := '202008091910_users';   
    DECLARE usr text := '';
    DECLARE admin_usr text := '';
    BEGIN  
        IF NOT (SELECT check_migration(migrationId)) THEN
          ------------------------------------------------------------
            --migration script start
            EXECUTE  'EXECUTE get_dbUser ' INTO usr;            
            EXECUTE  'EXECUTE get_dbAdmin ' INTO admin_usr;
            ------------------------------------------------------------
            -- SCHEMA:
            CREATE TABLE IF NOT EXISTS users (
                id serial PRIMARY KEY,
                name CHARACTER VARYING NOT NULL,
                role_id BIGINT NOT NULL,
                UNIQUE(name)
            );                            
                
            EXECUTE 'ALTER TABLE deputy.roles OWNER TO ' || quote_ident(admin_usr);
            EXECUTE 'GRANT ALL ON TABLE deputy.roles TO ' || quote_ident(usr); 

            CREATE INDEX ON users(role_id);
            ALTER TABLE users
                ADD CONSTRAINT user_role_fk FOREIGN KEY (role_id) REFERENCES roles(id);

            ---------------------------------------------------------------------------------------------------
            -- This function returuns all required buckets for the requested period. Each row represents one bucket for one metric (action/category)
            -- plus one NULL bucket for each metric (action/category) which represents the previous period total
            DROP PROCEDURE IF EXISTS setUsers();
            
            CREATE OR REPLACE PROCEDURE setUsers()
            LANGUAGE plpgsql

            AS
            $pr$
            BEGIN
                INSERT INTO users (name, role_id)
                VALUES
                    ('Adam Admin', 1),
                    ('Emily Employee', 4),
                    ('Sam Supervisor', 3),
                    ('Mary Manager', 2),
                    ('Steve Trainer', 5)
                ON CONFLICT ON CONSTRAINT users_name_key DO NOTHING;

                COMMIT;
            END;
            $pr$;
            
            ------------------------------------------------------------  
            --migration script end
            PERFORM create_migration(migrationId);
            ------------------------------------------------------------
        END IF;
    END    
    $$;
COMMIT;
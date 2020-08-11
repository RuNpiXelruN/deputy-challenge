\set ON_ERROR_STOP on

BEGIN;
    DO $$
    -- DECLARE migrationId VARCHAR := '202008091900_roles';   
    -- DECLARE usr text := '';
    -- DECLARE admin_usr text := '';
        -- IF NOT (SELECT check_migration(migrationId)) THEN
          ------------------------------------------------------------
            --migration script start
            -- EXECUTE  'EXECUTE get_dbUser ' INTO usr;            
            -- EXECUTE  'EXECUTE get_dbAdmin ' INTO admin_usr;
            ------------------------------------------------------------
            -- SCHEMA:
            CREATE TABLE IF NOT EXISTS roles (
                id serial PRIMARY KEY,
                name CHARACTER VARYING NOT NULL,
                parent_id BIGINT,
                UNIQUE(name)
            );                            
                
            -- EXECUTE 'ALTER TABLE roles OWNER TO ' || quote_ident(admin_usr);
            -- EXECUTE 'GRANT ALL ON TABLE roles TO ' || quote_ident(usr); 

            ALTER TABLE roles
                ADD CONSTRAINT role_parent_fk FOREIGN KEY (parent_id) REFERENCES roles(id);

            ---------------------------------------------------------------------------------------------------
            -- This function returuns all required buckets for the requested period. Each row represents one bucket for one metric (action/category)
            -- plus one NULL bucket for each metric (action/category) which represents the previous period total 
            DROP PROCEDURE IF EXISTS setRoles();

            CREATE OR REPLACE PROCEDURE setRoles()
            LANGUAGE plpgsql

            AS
            $pr$
            BEGIN
                INSERT INTO roles (name, parent_id)
                VALUES
                    ('System Administrator', 0),
                    ('Location Manager', 1),
                    ('Supervisor', 2),
                    ('Employee', 3),
                    ('Trainer', 3)
                ON CONFLICT ON CONSTRAINT roles_name_key DO NOTHING;

                COMMIT;
            END;
            $pr$;

            ------------------------------------------------------------  
            --migration script end
            -- PERFORM create_migration(migrationId);
            ------------------------------------------------------------
        -- END IF;
    $$;
COMMIT;
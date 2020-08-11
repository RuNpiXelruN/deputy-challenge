\set ON_ERROR_STOP on

BEGIN;
    DO $$
    DECLARE migrationId VARCHAR := '202008091900_roles';   
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

            RAISE NOTICE 'Runninggggg';

            BEGIN
                CREATE TABLE deputy.roles (
                    id serial PRIMARY KEY,
                    name CHARACTER VARYING NOT NULL,
                    parent_id BIGINT
                );                            
                
            END;
            EXECUTE 'ALTER TABLE deputy.roles OWNER TO ' || quote_ident(admin_usr);
            EXECUTE 'GRANT ALL ON TABLE deputy.roles TO ' || quote_ident(usr); 

            CREATE INDEX ON deputy.roles(id);
            ALTER TABLE deputy.roles
                ADD CONSTRAINT role_parent_fk FOREIGN KEY (parent_id) REFERENCES deputy.roles(id);

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


-- BEGIN
-- 	CREATE TABLE roles (
-- 		id serial PRIMARY KEY,
-- 		name CHARACTER VARYING NOT NULL,
-- 		parent_id BIGINT
-- 	);
	
-- 	CREATE TABLE users (
-- 		id serial PRIMARY KEY,
-- 		name CHARACTER VARYING NOT NULL,
-- 		role_id BIGINT NOT NULL
-- 	);
	
-- 	CREATE INDEX ON roles(id);
	
-- 	CREATE INDEX ON users(role_id);
	
-- 	ALTER TABLE roles
-- 		ADD CONSTRAINT role_parent_fk FOREIGN KEY parent_id REFERENCES roles(id);
-- 	ALTER TABLE users
-- 		ADD CONSTRAINT user_role_fk FOREIGN KEY role_id REFERENCES roles(id);
-- END;
/*Parameters in this file::
	v0=databasename
	v2=owneruser
	v1=username,
    v4=password

    dbName={database name}
    dbUser={(service) user name}
    dbAdmin={admin user}
    dbPass={'password for (service) user'}
*/
-- this executes outside of transaction blocks.
\i config.sql
CREATE DATABASE :dbName
    WITH
    OWNER = :dbAdmin
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    CONNECTION LIMIT = -1;

\c :dbName;
-- You have to create the DB first and then get a connection to the DB you created, then you can run this :
\set ON_ERROR_STOP on
BEGIN;
    -- this makes the prepared statements available for all the transactional code blocks in subsequent files
    \i config.sql
    --schema might exist already
    CREATE SCHEMA IF NOT EXISTS deputy;

    GRANT ALL ON SCHEMA deputy TO :dbAdmin;

    DO $$
    DECLARE usr text := '';
    DECLARE pwd text := '';
    BEGIN

        EXECUTE  'EXECUTE get_dbUser ' INTO usr;
        EXECUTE  'EXECUTE get_dbPass ' INTO pwd;

        IF NOT EXISTS (SELECT
            FROM   pg_catalog.pg_roles
            WHERE  rolname = usr) THEN
                execute 'CREATE ROLE ' ||  usr || ' WITH LOGIN ENCRYPTED PASSWORD '  ||' ''' || pwd  || ''' ';
        END IF;
    END;
    $$;

	GRANT ALL ON SCHEMA deputy TO :dbUser;

    --since this is the setup script, migrations shouldn't exist.
    --However adding if not exists just to be safe in case some one created the table manually
    CREATE TABLE IF NOT EXISTS deputy.migrations (
        migration_id varchar(100) PRIMARY KEY NOT NULL,
        date_migration_ran timestamp DEFAULT now() NOT NULL
    );

    CREATE OR REPLACE FUNCTION deputy.check_migration(migrationId varchar(100))
    RETURNS boolean

    AS
    $function$
		DECLARE migrationExists boolean;
        BEGIN
            IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = 'deputy' AND table_name = 'migrations') THEN

                migrationExists := EXISTS(SELECT 1 FROM curator.migrations WHERE migration_id = migrationId);

				IF (migrationExists) THEN
					RAISE NOTICE 'MIGRATION % ALREADY EXISTS', migrationId;
				END IF;

				RETURN migrationExists;

            ELSE
                RAISE EXCEPTION 'MIGRATIONS TABLE DOES NOT EXIST';
            END IF;
        END;
    $function$;
    LANGUAGE 'plpgsql'


    CREATE OR REPLACE FUNCTION deputy.create_migration(migrationId varchar(100))
    RETURNS void
        
    AS
    $function$
        BEGIN
            IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = 'deputy' AND table_name = 'migrations') THEN
                INSERT INTO curator.migrations (migration_id) VALUES (migrationId);
            ELSE
                RAISE EXCEPTION 'MIGRATIONS TABLE DOES NOT EXIST';
            END IF;
        END;
    $function$;
    LANGUAGE 'plpgsql'

COMMIT;
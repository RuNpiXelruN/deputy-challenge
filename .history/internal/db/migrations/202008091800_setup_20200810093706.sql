/*Parameters in this file::
	v0=databasename
	v2=owneruser
	v1=username,
    v4=password
*/
-- this executes outside of transaction blocks.
\i config.sql
CREATE DATABASE :v0
    WITH
    OWNER = :v2
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    CONNECTION LIMIT = -1;

\c :v0;
-- You have to create the DB first and then get a connection to the DB you created, then you can run this :
\set ON_ERROR_STOP on
BEGIN;
    -- this makes the prepared statements available for all the transactional code blocks in subsequent files
    \i config.sql
    --schema might exist already
    CREATE SCHEMA IF NOT EXISTS deputy;

    GRANT ALL ON SCHEMA deputy TO :v2;

    DO $$
    DECLARE usr text := '';
    DECLARE pwd text := '';
    BEGIN

        EXECUTE  'EXECUTE get_v1 ' INTO usr;
        EXECUTE  'EXECUTE get_v4 ' INTO pwd;

        IF NOT EXISTS (SELECT
            FROM   pg_catalog.pg_roles
            WHERE  rolname = usr) THEN
                execute 'CREATE ROLE ' ||  usr || ' WITH LOGIN ENCRYPTED PASSWORD '  ||' ''' || pwd  || ''' ';
        END IF;
    END;
    $$;

	GRANT ALL ON SCHEMA deputy TO :v1;

    --since this is the setup script, migrations shouldn't exist.
    --However adding if not exists just to be safe in case some one created the table manually
    CREATE TABLE IF NOT EXISTS curator.migrations (
        migration_id varchar(100) PRIMARY KEY NOT NULL,
        date_migration_ran timestamp DEFAULT now() NOT NULL
    );

    CREATE OR REPLACE FUNCTION curator.check_migration(migrationId varchar(100))
    RETURNS boolean
    LANGUAGE 'plpgsql'
    AS
    $func$
		DECLARE migrationExists boolean;
        BEGIN
            IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = 'curator' AND table_name = 'migrations') THEN

                migrationExists := EXISTS(SELECT 1 FROM curator.migrations WHERE migration_id = migrationId);

				IF (migrationExists) THEN
					RAISE NOTICE 'MIGRATION % ALREADY EXISTS', migrationId;
				END IF;

				RETURN migrationExists;

            ELSE
                RAISE EXCEPTION 'MIGRATIONS TABLE DOES NOT EXIST';
            END IF;
        END;
    $func$;

    CREATE OR REPLACE FUNCTION curator.create_migration(migrationId varchar(100))
    RETURNS void
    LANGUAGE 'plpgsql'
    AS
    $func$
        BEGIN
            IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = 'curator' AND table_name = 'migrations') THEN
                INSERT INTO curator.migrations (migration_id) VALUES (migrationId);
            ELSE
                RAISE EXCEPTION 'MIGRATIONS TABLE DOES NOT EXIST';
            END IF;
        END;
    $func$;

COMMIT;
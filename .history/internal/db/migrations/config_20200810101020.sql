--\set v0 {database name}
--\set v1 {(service) user name}
--\set v2 {admin user}
--\set v4 {'password for (service) user'}


-- Local:
\set dbName deputychallenge
\set dbUser deputyuser
\set dbAdmin deputyadmin
\set dbPass 'password'


PREPARE get_dbName AS SELECT * FROM regexp_split_to_table(:'dbName', ',');
PREPARE get_dbUser AS SELECT * FROM regexp_split_to_table(:'dbUser', ',');
PREPARE get_dbAdmin AS SELECT * FROM regexp_split_to_table(:'dbAdmin', ',');
PREPARE get_dbPass AS SELECT * FROM regexp_split_to_table(:'dbPass', ',');

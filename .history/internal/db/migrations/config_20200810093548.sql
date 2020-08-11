--\set v0 {database name}
--\set v1 {(service) user name}
--\set v2 {admin user}
--\set v4 {'password for (service) user'}


-- Local:
\set v0 deputychallenge
\set v1 postgres
\set v2 deputyuser
\set v4 'password'


PREPARE get_v1 AS SELECT * FROM regexp_split_to_table(:'v1', ',');
PREPARE get_v2 AS SELECT * FROM regexp_split_to_table(:'v2', ',');
PREPARE get_v4 AS SELECT * FROM regexp_split_to_table(:'v4', ',');
PREPARE get_v0 AS SELECT * FROM regexp_split_to_table(:'v0', ',');

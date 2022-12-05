CREATE EXTENSION if not exists pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE SCHEMA if not exists todo;

\i ./v001_crate_table.sql
\i ./v002_migrate_usr.sql

/*
  to run script use:
  psql -U postgres -d postgres -a -f .\create_db.sql
  or run create db task in vscode
*/

CREATE DATABASE whatsappmessaging
      WITH
      OWNER = postgres
      TEMPLATE = template0
      ENCODING = 'UTF8'
      LC_COLLATE = 'fa_IR.UTF8'
      LC_CTYPE = 'fa_IR.UTF8'
      TABLESPACE = pg_default
      CONNECTION LIMIT = -1;

ALTER DATABASE whatsappmessaging SET default_transaction_isolation TO 'read committed';

CREATE ROLE whatsappmessaging_services WITH
    LOGIN
    NOSUPERUSER
    INHERIT
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    PASSWORD 'services';

ALTER Role whatsappmessaging_services SET default_transaction_isolation TO 'read committed';
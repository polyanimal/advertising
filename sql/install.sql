\set ON_ERROR_STOP 1

DROP DATABASE IF EXISTS mdb;
DROP user IF EXISTS mdb;
CREATE DATABASE mdb;
CREATE user mdb WITH PASSWORD 'mdb';

\connect mdb

CREATE SCHEMA mdb;
GRANT usage ON SCHEMA mdb TO mdb;

create table mdb.advertisement
(
    id varchar(100) PRIMARY KEY,
    name varchar(200) NOT NULL ,
    description varchar(1000),
    photo_links varchar(100)[],
    price integer NOT NULL,
    date_create timestamp NOT NULL DEFAULT NOW()
);
GRANT SELECT, INSERT, UPDATE, DELETE ON mdb.advertisement TO mdb;
COMMENT ON TABLE mdb.advertisement IS 'Объявление';
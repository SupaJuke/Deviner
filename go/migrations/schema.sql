DROP SCHEMA IF EXISTS :DB_SCHEMA CASCADE;
CREATE SCHEMA :DB_SCHEMA;
SET search_path TO :DB_SCHEMA;

CREATE TABLE Users(
	username VARCHAR(255) PRIMARY KEY,
	password TEXT,
	code 	 TEXT DEFAULT '63556' -- Nello
);

INSERT INTO Users (username, password)
VALUES
('admin', '');
-- ('test', '');

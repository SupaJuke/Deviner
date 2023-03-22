DROP SCHEMA IF EXISTS pooe_game CASCADE;
CREATE SCHEMA pooe_game;
SET search_path TO pooe_game;

CREATE TABLE Users(
	username VARCHAR(255) PRIMARY KEY,
	password TEXT
);

INSERT INTO Users (username, password)
VALUES
('admin', ''),
('test', '');

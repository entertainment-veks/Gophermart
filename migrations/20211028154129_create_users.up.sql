CREATE TABLE IF NOT EXISTS users (
    id bigserial NOT NULL PRIMARY KEY, 
    login text NOT NULL UNIQUE,
    password text NOT NULL
);
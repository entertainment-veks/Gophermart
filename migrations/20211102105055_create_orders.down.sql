CREATE TABLE IF NOT EXISTS orders (
    id bigserial NOT NULL PRIMARY KEY, 
    number bigserial NOT NULL UNIQUE,
    status text NOT NULL,
    owner text NOT NULL
);
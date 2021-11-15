CREATE TABLE IF NOT EXISTS operations (
    id bigserial NOT NULL PRIMARY KEY, 
    owner text NOT NULL,
    order_num bigserial NOT NULL,
    amount int NOT NULL,
    processed_at timestamptz NOT NULL
);
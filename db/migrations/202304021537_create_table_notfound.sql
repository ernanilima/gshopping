-- +goose Up
CREATE TABLE IF NOT EXISTS notfound (
    id serial PRIMARY KEY,
    barcode VARCHAR(14) NOT NULL UNIQUE,
    attempts int4 NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS notfound;

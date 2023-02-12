-- +goose Up
CREATE TABLE IF NOT EXISTS brand (
    id UUID PRIMARY KEY,
    description VARCHAR(50) NOT NULL UNIQUE,
    created_date TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS brand;

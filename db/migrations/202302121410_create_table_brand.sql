-- +goose Up
CREATE TABLE IF NOT EXISTS brand (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code SERIAL NOT NULL UNIQUE,
    description VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS brand;

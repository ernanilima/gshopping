-- +goose Up
CREATE EXTENSION IF NOT EXISTS "unaccent" schema pg_catalog;

-- +goose Down
DROP EXTENSION IF EXISTS "unaccent" schema pg_catalog;

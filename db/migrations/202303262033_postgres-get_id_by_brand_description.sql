-- +goose Up
CREATE OR REPLACE FUNCTION get_brand_id_by_description(v VARCHAR) RETURNS UUID AS $$ BEGIN RETURN id FROM brand WHERE UPPER(unaccent(description)) = UPPER(unaccent(v)); END; $$ LANGUAGE plpgsql;

-- +goose Down
DROP FUNCTION IF EXISTS get_brand_id_by_description;

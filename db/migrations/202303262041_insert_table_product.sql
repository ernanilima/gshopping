-- +goose Up
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('789102030', 'Produto De Teste', get_brand_id_by_description('1906'), LOCALTIMESTAMP);

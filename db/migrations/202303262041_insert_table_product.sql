-- +goose Up
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('789102030', 'Produto De Teste', get_brand_id_by_description('1906'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('8412598001079', 'Cerveja Vienna Lager Red Vintage 1906 Ga', get_brand_id_by_description('1906'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005802892', 'Cappuccino Solúvel Chocolate 3 Corações ', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005800676', 'Café Solúvel Granulado Tradicional 3 Cor', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005800690', 'Café Solúvel Granulado Descafeinado 3 Co', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005800683', 'Café Solúvel Granulado Tradicional 3 Cor', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005800706', 'Café Solúvel Granulado Descafeinado 3 Co', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005802403', 'Cappuccino Solúvel Baunilha 3 Corações S', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005802427', 'Cappuccino Solúvel Canela 3 Corações Sac', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005802397', 'Cappuccino Solúvel Chocolate 3 Corações ', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
INSERT INTO product (barcode, description, brand_id, created_at) VALUES ('7896005800157', 'Cappuccino Solúvel Classic 3 Corações Po', get_brand_id_by_description('3 Corações'), LOCALTIMESTAMP);
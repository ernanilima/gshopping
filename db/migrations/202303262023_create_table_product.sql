-- +goose Up
CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    barcode VARCHAR(14) NOT NULL UNIQUE,
    description VARCHAR(50) NOT NULL UNIQUE,
    brand_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_product_brand FOREIGN KEY (brand_id) REFERENCES public.brand(id)
);

-- +goose Down
DROP TABLE IF EXISTS product;

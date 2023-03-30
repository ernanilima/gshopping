package product_repository

import (
	"github.com/ernanilima/gshopping/app/model"
)

func (c *ProductConnection) FindByBarcode(barcode string) (model.Product, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow(`
		SELECT p.id id, p.barcode barcode, p.description description, b.description brand, p.created_at created_at FROM product p
			JOIN Brand b ON b.id = p.brand_id
			WHERE barcode = $1`, barcode)

	var product model.Product
	if err := result.Scan(&product.ID, &product.Barcode, &product.Description, &product.Brand, &product.CreatedAt); err != nil {
		return model.Product{}, err
	}

	return product, nil
}

package product_repository

import (
	"log"

	"github.com/ernanilima/gshopping/app/model"
)

// FindByBarcode busca um produto pelo codigo de barras
func (c *ProductConnection) FindByBarcode(barcode string) (model.Product, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow(`
		SELECT p.id id, p.barcode barcode, p.description description, b.description brand, p.created_at created_at FROM product p
			JOIN Brand b ON b.id = p.brand_id
			WHERE barcode = $1`, barcode)

	var product model.Product
	if err := result.Scan(&product.ID, &product.Barcode, &product.Description, &product.Brand, &product.CreatedAt); err != nil {
		c.notFound(barcode)
		return model.Product{}, err
	}

	return product, nil
}

// notFound registra um produto nao localizado por codigo de barras
func (c *ProductConnection) notFound(barcode string) {
	conn := c.OpenConnection()
	defer conn.Close()

	_, err := conn.Exec(`
		INSERT INTO notfound (barcode, attempts) VALUES ($1, 1) ON CONFLICT (barcode)
			DO UPDATE SET attempts = notfound.attempts + 1`, barcode)
	if err != nil {
		log.Fatalf("Erro ao inserir o produto com o codigo de barras: %s", barcode)
	}
}

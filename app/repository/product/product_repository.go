package product_repository

import (
	"fmt"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
)

// FindAllProducts busca uma lista paginada de produtos com base em um filtro opcional
func (c *ProductConnection) FindAllProducts(filter string, pageable utils.Pageable) (utils.Pageable, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), p.id id, p.barcode barcode, p.description description, b.description brand, p.created_at created_at FROM product p
			JOIN brand b ON b.id = p.brand_id
			WHERE (CASE 
				WHEN $1 = '' THEN TRUE 
				ELSE (p.barcode ILIKE $1
					OR UPPER(unaccent(p.description)) ILIKE $1
					OR UPPER(unaccent(b.description)) ILIKE $1)
			END)
			ORDER BY %s
			LIMIT $2 OFFSET $3`, pageable.Sort)

	results, err := conn.Query(query, fmt.Sprintf("%%%s%%", filter), pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		return utils.Pageable{}, err
	}
	defer results.Close()

	var products []model.Product
	for results.Next() {
		var product model.Product
		results.Scan(&pageable.TotalElements, &product.ID, &product.Barcode, &product.Description, &product.Brand, &product.CreatedAt)
		products = append(products, product)
	}

	return utils.GeneratePaginationRequest(products, pageable), nil
}

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
		fmt.Printf("Erro ao inserir o produto com o codigo de barras: %s | %s", barcode, err)
	}
}

// FindAllNotFound busca uma lista com todos os produtos nao encontrados
func (c *ProductConnection) FindAllNotFound(pageable utils.Pageable) utils.Pageable {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), * FROM notfound nf
			ORDER BY %s
			LIMIT $1 OFFSET $2`, pageable.Sort)

	results, err := conn.Query(query, pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		return utils.Pageable{}
	}
	defer results.Close()

	var productsNotFound []model.ProductNotFound
	for results.Next() {
		var productNotFound model.ProductNotFound
		results.Scan(&pageable.TotalElements, &productNotFound.ID, &productNotFound.Barcode, &productNotFound.Attempts)
		productsNotFound = append(productsNotFound, productNotFound)
	}

	return utils.GeneratePaginationRequest(productsNotFound, pageable)
}

func (c *ProductConnection) FindNotFoundByBarcode(barcode string, pageable utils.Pageable) (utils.Pageable, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), * FROM notfound nf
			WHERE nf.barcode ILIKE $1
			ORDER BY %s
			LIMIT $2 OFFSET $3`, pageable.Sort)

	results, err := conn.Query(query, fmt.Sprintf("%%%s%%", barcode), pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		return utils.Pageable{}, err
	}
	defer results.Close()

	var productsNotFound []model.ProductNotFound
	for results.Next() {
		var productNotFound model.ProductNotFound
		results.Scan(&pageable.TotalElements, &productNotFound.ID, &productNotFound.Barcode, &productNotFound.Attempts)
		productsNotFound = append(productsNotFound, productNotFound)
	}

	return utils.GeneratePaginationRequest(productsNotFound, pageable), nil
}

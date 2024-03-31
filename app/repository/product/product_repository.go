package product_repository

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
)

// InsertProduct insere um produto
func (c *ProductConnection) InsertProduct(product model.Product) (model.Product, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	product.ID = uuid.New()
	product.CreatedAt = time.Now()

	_, err := conn.Exec("INSERT INTO product (id, barcode, description, brand_id, created_at) VALUES ($1, $2, $3, $4, $5)",
		product.ID, strings.TrimSpace(product.Barcode), strings.TrimSpace(product.Description), product.Brand.ID, product.CreatedAt)
	if err != nil {
		print("InsertProduct - ", err.Error())
		return model.Product{}, err
	}

	c.deleteProductNotFound(product.Barcode)

	return product, nil
}

// EditProduct edita um produto
func (c *ProductConnection) EditProduct(product model.Product) (model.Product, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result, _ := c.FindProductById(product.ID)
	result.Barcode = product.Barcode
	result.Description = product.Description
	result.Brand.ID = product.Brand.ID

	_, err := conn.Exec("UPDATE product SET barcode=$2, description=$3, brand_id=$4 WHERE id=$1",
		result.ID, strings.TrimSpace(result.Barcode), strings.TrimSpace(result.Description), result.Brand.ID)
	if err != nil {
		print("EditProduct - ", err.Error())
		return model.Product{}, err
	}

	c.deleteProductNotFound(result.Barcode)

	return result, nil
}

// FindAllProducts busca uma lista paginada de produtos com base em um filtro opcional
func (c *ProductConnection) FindAllProducts(filter string, pageable utils.Pageable) utils.Pageable {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		WITH produtos AS (
			SELECT p.id id, 
				p.barcode barcode,
				p.description description,
				b.id brand_id,
				b.description brand,
				p.created_at created_at
			FROM product p
			JOIN brand b ON b.id = p.brand_id
			WHERE (CASE 
				WHEN $1 = '' THEN TRUE 
				ELSE (p.barcode ILIKE $1
					OR UPPER(unaccent(p.description)) ILIKE $1
					OR UPPER(unaccent(b.description)) ILIKE $1)
			END)
		),
		total_products AS (
			SELECT COUNT(*) total
			FROM produtos
		)
		SELECT total, p.*
			FROM produtos p
			CROSS JOIN total_products total
			ORDER BY %s
			LIMIT $2 OFFSET $3`, pageable.Sort)

	results, err := conn.Query(query, fmt.Sprintf("%%%s%%", filter), pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		print("FindAllProducts - ", err.Error())
		return utils.Pageable{}
	}
	defer results.Close()

	var products []model.Product
	for results.Next() {
		var product model.Product
		results.Scan(&pageable.TotalElements, &product.ID, &product.Barcode, &product.Description, &product.Brand.ID, &product.Brand.Description, &product.CreatedAt)
		products = append(products, product)
	}

	return utils.GeneratePaginationRequest(products, pageable)
}

// FindProductById busca um produto pelo ID
func (c *ProductConnection) FindProductById(id uuid.UUID) (model.Product, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow(`
		SELECT p.id id,
				p.barcode barcode,
				p.description description,
				p.created_at created_at,
				b.id brand_id,
				b.description brand_description
			FROM product p
			LEFT JOIN brand b ON b.id = p.brand_id
			WHERE p.id = $1`, id)

	var product model.Product
	if err := result.Scan(&product.ID, &product.Barcode, &product.Description, &product.CreatedAt, &product.Brand.ID, &product.Brand.Description); err != nil {
		print("FindProductById - ", err.Error())
		return model.Product{}, err
	}

	return product, nil
}

// FindProductByBarcode busca um produto pelo codigo de barras
func (c *ProductConnection) FindProductByBarcode(barcode string) (model.Product, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	regex := regexp.MustCompile(`\D`)
	barcode = regex.ReplaceAllString(barcode, "")

	result := conn.QueryRow(`
		SELECT p.id id,
				p.barcode barcode,
				p.description description,
				b.description brand_description,
				p.created_at created_at
			FROM product p
			JOIN Brand b ON b.id = p.brand_id
			WHERE barcode = $1`, barcode)

	var product model.Product
	if err := result.Scan(&product.ID, &product.Barcode, &product.Description, &product.Brand.Description, &product.CreatedAt); err != nil {
		if len(barcode) > 0 {
			c.insertProductNotFound(barcode)
		}
		print("FindProductByBarcode - ", err.Error())
		return model.Product{}, err
	}

	return product, nil
}

// FindAllProductsNotFound busca uma lista com todos os produtos nao encontrados
func (c *ProductConnection) FindAllProductsNotFound(pageable utils.Pageable) utils.Pageable {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), * FROM notfound nf
			ORDER BY %s
			LIMIT $1 OFFSET $2`, pageable.Sort)

	results, err := conn.Query(query, pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		print("FindAllProductsNotFound - ", err.Error())
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

// FindAllProductsNotFoundByBarcode busca uma lista com todos os produtos nao encontrados pelo codigo de barras
func (c *ProductConnection) FindAllProductsNotFoundByBarcode(barcode string, pageable utils.Pageable) (utils.Pageable, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), * FROM notfound nf
			WHERE nf.barcode ILIKE $1
			ORDER BY %s
			LIMIT $2 OFFSET $3`, pageable.Sort)

	results, err := conn.Query(query, fmt.Sprintf("%%%s%%", barcode), pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		print("FindAllProductsNotFoundByBarcode - ", err.Error())
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

// insertProductNotFound registra um produto nao localizado por codigo de barras
func (c *ProductConnection) insertProductNotFound(barcode string) {
	conn := c.OpenConnection()
	defer conn.Close()

	_, err := conn.Exec(`
		INSERT INTO notfound (barcode, attempts) VALUES ($1, 1) ON CONFLICT (barcode)
			DO UPDATE SET attempts = notfound.attempts + 1`, barcode)
	if err != nil {
		fmt.Printf("Erro ao inserir ProductNotFound com o codigo de barras: %s | %s", barcode, err)
	}
}

// deleteProductNotFound deleta um produto nao localizado por codigo de barras
func (c *ProductConnection) deleteProductNotFound(barcode string) {
	conn := c.OpenConnection()
	defer conn.Close()

	_, err := conn.Exec(`
		DELETE FROM notfound WHERE barcode=$1`, barcode)
	if err != nil {
		fmt.Printf("Erro ao deletar ProductNotFound com o codigo de barras: %s | %s", barcode, err)
	}
}

// FindTotalProducts busca o total de produtos cadastrados
func (c *ProductConnection) FindTotalProducts() int32 {
	conn := c.OpenConnection()
	defer conn.Close()

	results := conn.QueryRow("SELECT COUNT(*) FROM product")

	var totalProducts int32
	if err := results.Scan(&totalProducts); err != nil {
		print("FindTotalProducts - ", err.Error())
		return 0
	}

	return totalProducts
}

// FindTotalProductsNotFound busca o total de produtos cadastrados
func (c *ProductConnection) FindTotalProductsNotFound() int32 {
	conn := c.OpenConnection()
	defer conn.Close()

	results := conn.QueryRow("SELECT COUNT(*) FROM notfound")

	var totalProductsNotFound int32
	if err := results.Scan(&totalProductsNotFound); err != nil {
		print("FindTotalProductsNotFound - ", err.Error())
		return 0
	}

	return totalProductsNotFound
}

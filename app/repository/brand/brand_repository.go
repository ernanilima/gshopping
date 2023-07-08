package brand_repository

import (
	"fmt"
	"time"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
)

// Insert insere uma marca
func (c *BrandConnection) Insert(brand model.Brand) (model.Brand, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	brand.ID = uuid.New()
	brand.CreatedAt = time.Now()

	_, err := conn.Exec("INSERT INTO brand (id, description, created_at) VALUES ($1, $2, $3)", brand.ID, brand.Description, brand.CreatedAt)
	if err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

// FindAll busca uma lista paginada de marcas
func (c *BrandConnection) FindAll(pageable utils.Pageable) utils.Pageable {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), COUNT(p.id) as total_products, b.* FROM brand b
			LEFT JOIN product p ON b.id = p.brand_id
			GROUP BY b.id
			ORDER BY %s
			LIMIT $1 OFFSET $2`, pageable.Sort)

	results, err := conn.Query(query, pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		return utils.Pageable{}
	}
	defer results.Close()

	var brands []model.Brand
	for results.Next() {
		var brand model.Brand
		results.Scan(&pageable.TotalElements, &brand.TotalProducts, &brand.ID, &brand.Code, &brand.Description, &brand.CreatedAt)
		brands = append(brands, brand)
	}

	return utils.GeneratePaginationRequest(brands, pageable)
}

// FindById busca uma marca pelo ID
func (c *BrandConnection) FindById(id uuid.UUID) (model.Brand, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow(`
		SELECT COUNT(p.id), b.* FROM brand b
			LEFT JOIN product p ON b.id = p.brand_id
			WHERE b.id = $1
			GROUP BY b.id`, id)

	var brand model.Brand
	if err := result.Scan(&brand.TotalProducts, &brand.ID, &brand.Code, &brand.Description, &brand.CreatedAt); err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

// FindByDescription busca uma lista paginada de marcas pela %descricao%
func (c *BrandConnection) FindByDescription(description string, pageable utils.Pageable) (utils.Pageable, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), COUNT(p.id) as total_products, b.* FROM brand b
			LEFT JOIN product p ON b.id = p.brand_id
			WHERE UPPER(unaccent(b.description)) ILIKE $1
			GROUP BY b.id
			ORDER BY %s
			LIMIT $2 OFFSET $3`, pageable.Sort)

	results, err := conn.Query(query, fmt.Sprintf("%%%s%%", description), pageable.Size, pageable.Size*pageable.Page)

	if err != nil {
		return utils.Pageable{}, err
	}
	defer results.Close()

	var brands []model.Brand
	for results.Next() {
		var brand model.Brand
		results.Scan(&pageable.TotalElements, &brand.TotalProducts, &brand.ID, &brand.Code, &brand.Description, &brand.CreatedAt)
		brands = append(brands, brand)
	}

	return utils.GeneratePaginationRequest(brands, pageable), nil
}

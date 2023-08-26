package brand_repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
)

// InsertBrand insere uma marca
func (c *BrandConnection) InsertBrand(brand model.Brand) (model.Brand, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	brand.ID = uuid.New()
	brand.CreatedAt = time.Now()

	_, err := conn.Exec("INSERT INTO brand (id, description, created_at) VALUES ($1, $2, $3)",
		brand.ID, strings.TrimSpace(brand.Description), brand.CreatedAt)
	if err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

// EditBrand edita uma marca
func (c *BrandConnection) EditBrand(brand model.Brand) (model.Brand, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result, _ := c.FindBrandById(brand.ID)
	result.Description = brand.Description

	_, err := conn.Exec("UPDATE brand SET description=$2 WHERE id=$1",
		result.ID, strings.TrimSpace(result.Description))
	if err != nil {
		return model.Brand{}, err
	}

	return result, nil
}

// DeleteBrand deleta uma marca
func (c *BrandConnection) DeleteBrand(id uuid.UUID) (model.Brand, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result, err := c.FindBrandById(id)
	if err != nil {
		return model.Brand{}, err
	}
	_, err = conn.Exec("DELETE FROM brand WHERE id=$1", id)
	if err != nil {
		return model.Brand{}, err
	}

	return result, nil
}

// FindAllBrands busca uma lista paginada de marcas
func (c *BrandConnection) FindAllBrands(pageable utils.Pageable) utils.Pageable {
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

// FindBrandById busca uma marca pelo ID
func (c *BrandConnection) FindBrandById(id uuid.UUID) (model.Brand, error) {
	conn := c.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow(`
		SELECT COUNT(p.id) as total_products, b.* FROM brand b
			LEFT JOIN product p ON b.id = p.brand_id
			WHERE b.id = $1
			GROUP BY b.id`, id)

	var brand model.Brand
	if err := result.Scan(&brand.TotalProducts, &brand.ID, &brand.Code, &brand.Description, &brand.CreatedAt); err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

// FindAllBrandsByDescription busca uma lista paginada de marcas pela %descricao%
func (c *BrandConnection) FindAllBrandsByDescription(description string, pageable utils.Pageable) (utils.Pageable, error) {
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

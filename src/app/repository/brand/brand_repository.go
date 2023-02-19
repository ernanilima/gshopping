package brand

import (
	"fmt"

	"github.com/ernanilima/gshopping/src/app/model"
	"github.com/ernanilima/gshopping/src/app/repository"
	"github.com/ernanilima/gshopping/src/app/utils"
	"github.com/google/uuid"
)

// FindAll busca uma lista paginada de marcas
func FindAll(pageable utils.Pageable) utils.Pageable {
	conn, _ := repository.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), * FROM brand
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
		results.Scan(&pageable.TotalElements, &brand.ID, &brand.Description, &brand.CreatedAt)
		brands = append(brands, brand)
	}

	return utils.GeneratePaginationRequest(brands, pageable)
}

// FindById busca uma marca pelo ID
func FindById(id uuid.UUID) (model.Brand, error) {
	conn, _ := repository.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow("SELECT * FROM brand WHERE id = $1", id)

	var brand model.Brand
	if err := result.Scan(&brand.ID, &brand.Description, &brand.CreatedAt); err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

// FindByDescription busca uma lista paginada de marcas pela %descricao%
func FindByDescription(description string, pageable utils.Pageable) (utils.Pageable, error) {
	conn, _ := repository.OpenConnection()
	defer conn.Close()

	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), * FROM brand
			WHERE description ILIKE $1
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
		results.Scan(&pageable.TotalElements, &brand.ID, &brand.Description, &brand.CreatedAt)
		brands = append(brands, brand)
	}

	return utils.GeneratePaginationRequest(brands, pageable), nil
}

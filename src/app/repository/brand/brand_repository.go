package brand

import (
	"fmt"

	"github.com/ernanilima/gshopping/src/app/model"
	"github.com/ernanilima/gshopping/src/app/repository"
	"github.com/ernanilima/gshopping/src/app/utils"
	"github.com/google/uuid"
)

// FindAll busca uma lista com todas as marcas
// limitando a 30 registros
func FindAll(pageable utils.Pageable) utils.Pageable {
	conn, _ := repository.OpenConnection()
	defer conn.Close()

	results, err := conn.Query(`
		SELECT COUNT(*) OVER(), * FROM brand
		ORDER BY "description"
		LIMIT $1 OFFSET $2`, pageable.Size, pageable.Size*pageable.Page)

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

// FindByDescription busca uma lista de marcas pela %descricao%
// limitando a 30 registros
func FindByDescription(description string) ([]model.Brand, error) {
	conn, _ := repository.OpenConnection()
	defer conn.Close()

	results, err := conn.Query("SELECT * FROM brand WHERE description ILIKE $1 LIMIT 30", fmt.Sprintf("%%%s%%", description))

	if err != nil {
		return nil, err
	}
	defer results.Close()

	var brands []model.Brand
	for results.Next() {
		var brand model.Brand
		results.Scan(&brand.ID, &brand.Description, &brand.CreatedAt)
		brands = append(brands, brand)
	}

	return brands, nil
}

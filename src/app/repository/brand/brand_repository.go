package brand

import (
	"github.com/ernanilima/gshopping/src/app/model"
	"github.com/ernanilima/gshopping/src/app/repository"
	"github.com/google/uuid"
)

// FindById busca uma marca pelo ID
func FindById(id uuid.UUID) (model.Brand, error) {
	conn, _ := repository.OpenConnection()
	defer conn.Close()

	result := conn.QueryRow("SELECT * FROM brand WHERE id = $1", id)

	var brand model.Brand
	if err := result.Scan(&brand.ID, &brand.Description, &brand.CreatedDate); err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

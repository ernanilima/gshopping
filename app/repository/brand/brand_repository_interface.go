package brand_repository

import (
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
)

func NewBrandRepository(connector database.DatabaseConnector) BrandRepository {
	return &BrandConnection{connector}
}

type BrandConnection struct {
	database.DatabaseConnector
}

type BrandRepository interface {
	Insert(model.Brand) (model.Brand, error)
	FindAll(pageable utils.Pageable) utils.Pageable
	FindById(id uuid.UUID) (model.Brand, error)
	FindByDescription(description string, pageable utils.Pageable) (utils.Pageable, error)
}

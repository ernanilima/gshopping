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
	InsertBrand(model.Brand) (model.Brand, error)
	EditBrand(model.Brand) (model.Brand, error)
	DeleteBrand(id uuid.UUID) (model.Brand, error)
	FindAllBrands(pageable utils.Pageable) utils.Pageable
	FindBrandById(id uuid.UUID) (model.Brand, error)
	FindAllBrandsByDescription(description string, pageable utils.Pageable) (utils.Pageable, error)
	FindTotalBrands() int32
}

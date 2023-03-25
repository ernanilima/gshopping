package repository

import (
	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
	"github.com/ernanilima/gshopping/app/repository/database"
)

func NewRepository(connector database.DatabaseConnector) Repository {
	return &repository{
		brand_repository.NewBrandRepository(connector),
	}
}

type Repository interface {
	brand_repository.BrandRepository
}

type repository struct {
	brand_repository.BrandRepository
}

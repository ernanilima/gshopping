package repository

import (
	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
	"github.com/ernanilima/gshopping/app/repository/database"
	product_repository "github.com/ernanilima/gshopping/app/repository/product"
)

func NewRepository(connector database.DatabaseConnector) Repository {
	return &repository{
		product_repository.NewProductRepository(connector),
		brand_repository.NewBrandRepository(connector),
	}
}

type Repository interface {
	product_repository.ProductRepository
	brand_repository.BrandRepository
}

type repository struct {
	product_repository.ProductRepository
	brand_repository.BrandRepository
}

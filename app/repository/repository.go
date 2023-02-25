package repository

import (
	"github.com/ernanilima/gshopping/app/config"
	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
)

func NewRepository(cfg *config.Config) Repository {
	return &repository{
		brand_repository.NewBrandRepository(cfg),
	}
}

type Repository interface {
	brand_repository.BrandRepository
}

type repository struct {
	brand_repository.BrandRepository
}

package brand_repository

import (
	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
)

func NewBrandRepository(cfg *config.Config) BrandRepository {
	return &connection{cfg}
}

type connection struct {
	*config.Config
}

type BrandRepository interface {
	FindAll(pageable utils.Pageable) utils.Pageable
	FindById(id uuid.UUID) (model.Brand, error)
	FindByDescription(description string, pageable utils.Pageable) (utils.Pageable, error)
}

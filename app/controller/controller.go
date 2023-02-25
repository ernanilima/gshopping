package controller

import (
	brand_controller "github.com/ernanilima/gshopping/app/controller/brand"
	"github.com/ernanilima/gshopping/app/repository"
)

func NewController(repo repository.Repository) Controller {
	return &controller{
		brand_controller.NewBrandController(repo),
	}
}

type controller struct {
	brand_controller.BrandController
}

type Controller interface {
	brand_controller.BrandController
}

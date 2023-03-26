package controller

import (
	brand_controller "github.com/ernanilima/gshopping/app/controller/brand"
	product_controller "github.com/ernanilima/gshopping/app/controller/product"
	"github.com/ernanilima/gshopping/app/repository"
)

func NewController(repo repository.Repository) Controller {
	return &controller{
		product_controller.NewProductController(repo),
		brand_controller.NewBrandController(repo),
	}
}

type controller struct {
	product_controller.ProductController
	brand_controller.BrandController
}

type Controller interface {
	product_controller.ProductController
	brand_controller.BrandController
}

package brand_controller

import (
	"net/http"

	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
)

func NewBrandController(repo brand_repository.BrandRepository) BrandController {
	return &brandRepository{repo}
}

type brandRepository struct {
	brand_repository.BrandRepository
}

type BrandController interface {
	FindAllBrands(w http.ResponseWriter, r *http.Request)
	FindBrandById(w http.ResponseWriter, r *http.Request)
	FindAllBrandsByDescription(w http.ResponseWriter, r *http.Request)
}

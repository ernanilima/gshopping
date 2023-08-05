package product_controller

import (
	"net/http"

	product_repository "github.com/ernanilima/gshopping/app/repository/product"
)

func NewProductController(repo product_repository.ProductRepository) ProductController {
	return &productRepository{repo}
}

type productRepository struct {
	product_repository.ProductRepository
}

type ProductController interface {
	FindProductByBarcode(w http.ResponseWriter, r *http.Request)
	FindAllProductNotFound(w http.ResponseWriter, r *http.Request)
}

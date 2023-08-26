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
	InsertProduct(w http.ResponseWriter, r *http.Request)
	EditProduct(w http.ResponseWriter, r *http.Request)
	FindAllProducts(w http.ResponseWriter, r *http.Request)
	FindProductById(w http.ResponseWriter, r *http.Request)
	FindProductByBarcode(w http.ResponseWriter, r *http.Request)
	FindAllProductsNotFound(w http.ResponseWriter, r *http.Request)
	FindAllProductsNotFoundByBarcode(w http.ResponseWriter, r *http.Request)
}

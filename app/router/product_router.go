package router

import (
	"net/http"

	product_controller "github.com/ernanilima/gshopping/app/controller/product"
)

// productRouter recebe o controller e retorna o metodo que deve ser utilizado pela rota
// correspondente ao URI e ao HTTPMethod para ProductController
func productRouter(controller product_controller.ProductController) []Router {
	return []Router{
		{
			URI:        "/v1/produto",
			HTTPMethod: http.MethodPost,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.InsertProduct(w, r)
			},
		},
		{
			URI:        "/v1/produto/{id}",
			HTTPMethod: http.MethodPut,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.EditProduct(w, r)
			},
		},
		{
			URI:        "/v1/produto",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindAllProducts(w, r)
			},
		},
		{
			URI:        "/v1/produto/pesquisa/{filter}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindAllProducts(w, r)
			},
		},
		{
			URI:        "/v1/produto/{id}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindProductById(w, r)
			},
		},
		{
			URI:        "/v1/produto/codigo-barras/{barcode}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindProductByBarcode(w, r)
			},
		},
		{
			URI:        "/v1/produto/nao-encontrado",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindAllProductsNotFound(w, r)
			},
		},
		{
			URI:        "/v1/produto/nao-encontrado/{barcode}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindAllProductsNotFoundByBarcode(w, r)
			},
		},
	}
}

package router

import (
	"net/http"

	brand_controller "github.com/ernanilima/gshopping/app/controller/brand"
)

// brandRouter recebe o controller e retorna o metodo que deve ser utilizado pela rota
// correspondente ao URI e ao HTTPMethod para BrandController
func brandRouter(controller brand_controller.BrandController) []Router {
	return []Router{
		{
			URI:        "/v1/marca",
			HTTPMethod: http.MethodPost,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.InsertBrand(w, r)
			},
		},
		{
			URI:        "/v1/marca/{id}",
			HTTPMethod: http.MethodPut,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.EditBrand(w, r)
			},
		},
		{
			URI:        "/v1/marca/{id}",
			HTTPMethod: http.MethodDelete,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.DeleteBrand(w, r)
			},
		},
		{
			URI:        "/v1/marca",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindAllBrands(w, r)
			},
		},
		{
			URI:        "/v1/marca/{id}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindBrandById(w, r)
			},
		},
		{
			URI:        "/v1/marca/descricao/{description}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindAllBrandsByDescription(w, r)
			},
		},
	}
}

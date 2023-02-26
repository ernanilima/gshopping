package router

import (
	"net/http"

	brand_controller "github.com/ernanilima/gshopping/app/controller/brand"
)

// brandRouter recebe o controller e retorna a funcao correspondente ao URI e ao HTTPMethod para BrandController
func brandRouter(controller brand_controller.BrandController) []Router {
	return []Router{
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

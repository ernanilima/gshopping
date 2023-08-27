package router

import (
	"net/http"

	"github.com/ernanilima/gshopping/app/controller"
)

// dashboardRouter recebe o controller e retorna o metodo que deve ser utilizado pela rota
func dashboardRouter(controller controller.Controller) []Router {
	return []Router{
		{
			URI:        "/v1/dashboard/total-marcas",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindTotalBrands(w, r)
			},
		},
		{
			URI:        "/v1/dashboard/total-produtos",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindTotalProducts(w, r)
			},
		},
		{
			URI:        "/v1/dashboard/total-produtos-nao-encontrados",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindTotalProductsNotFound(w, r)
			},
		},
	}
}

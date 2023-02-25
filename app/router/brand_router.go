package router

import (
	"net/http"

	"github.com/ernanilima/gshopping/app/controller"
)

var brandRouter = []Router{
	{
		URI:        "/v1/marca",
		HTTPMethod: http.MethodGet,
		Function: func(c controller.Controller, w http.ResponseWriter, r *http.Request) {
			c.FindAllBrands(w, r)
		},
	},
	{
		URI:        "/v1/marca/{id}",
		HTTPMethod: http.MethodGet,
		Function: func(c controller.Controller, w http.ResponseWriter, r *http.Request) {
			c.FindBrandById(w, r)
		},
	},
	{
		URI:        "/v1/marca/descricao/{description}",
		HTTPMethod: http.MethodGet,
		Function: func(c controller.Controller, w http.ResponseWriter, r *http.Request) {
			c.FindAllBrandsByDescription(w, r)
		},
	},
}

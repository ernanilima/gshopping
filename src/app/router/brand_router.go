package router

import (
	"net/http"

	"github.com/ernanilima/gshopping/src/app/controller/brand"
)

var brandRouter = []Router{
	{
		URI:        "/v1/marca/{id}",
		HTTPMethod: http.MethodGet,
		Function:   brand.FindById,
	},
	{
		URI:        "/v1/marca/descricao/{description}",
		HTTPMethod: http.MethodGet,
		Function:   brand.FindByDescription,
	},
}

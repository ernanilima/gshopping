package router

import (
	"net/http"

	"github.com/ernanilima/gshopping/src/app/controller"
)

var productRouter = []Router{
	{
		URI:        "/v1/produto",
		HTTPMethod: http.MethodGet,
		Function:   controller.FindByBarcode,
	},
}

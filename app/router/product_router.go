package router

import (
	"net/http"

	"github.com/ernanilima/gshopping/app/controller/product"
)

var productRouter = []Router{
	{
		URI:        "/v1/produto",
		HTTPMethod: http.MethodGet,
		Function:   product.FindByBarcode,
	},
}

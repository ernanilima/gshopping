package router

import (
	"net/http"

	product_controller "github.com/ernanilima/gshopping/app/controller/product"
)

func productRouter(controller product_controller.ProductController) []Router {
	return []Router{
		{
			URI:        "/v1/produto/{barcode}",
			HTTPMethod: http.MethodGet,
			Function: func(w http.ResponseWriter, r *http.Request) {
				controller.FindProductByBarcode(w, r)
			},
		},
	}
}

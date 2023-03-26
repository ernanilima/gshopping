package router

import (
	"net/http"

	"github.com/ernanilima/gshopping/app/controller"
	"github.com/go-chi/chi"
)

// Router eh o tipo que sera usado para criar todas as rotas
type Router struct {
	URI        string
	HTTPMethod string
	Function   func(http.ResponseWriter, *http.Request)
}

// StartRoutes inicia as rotas
func StartRoutes(controller controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	routes := []Router{}
	routes = append(routes, productRouter(controller)...)
	routes = append(routes, brandRouter(controller)...)

	for _, router := range routes {
		r.Method(router.HTTPMethod, router.URI, http.HandlerFunc(router.Function))
	}

	return r
}

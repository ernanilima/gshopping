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
	Function   func(controller.Controller, http.ResponseWriter, *http.Request)
}

// StartRoutes inicia as rotas
func StartRoutes(c controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	routes := []Router{}
	routes = append(routes, productRouter...)
	routes = append(routes, brandRouter...)

	for _, router := range routes {
		r.Method(router.HTTPMethod, router.URI, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			router.Function(c, w, r)
		}))
	}

	return r
}

package router_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/ernanilima/gshopping/app/router"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Deve existir todas as rotas
func TestStartRoutes_Should_Exist_All_Routes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)

	r := router.StartRoutes(controller)

	routes := []http.Handler{
		getRouteByName(r.Routes(), "/v1/marca").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/marca/{id}").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/marca/descricao/{description}").Handlers["GET"],
	}

	assert.Equal(t, getTotalRoutesByName(r.Routes(), "/v1/marca"), len(routes), "Deve existir 3 rotas")
	for _, route := range routes {
		assert.NotNil(t, route, "deve ter o metodo GET")
	}
}

func getTotalRoutesByName(routes []chi.Route, route string) int {
	var total = 0

	for _, r := range routes {
		if strings.Contains(r.Pattern, route) {
			total++
		}
	}

	return total
}

func getRouteByName(routes []chi.Route, route string) chi.Route {
	var matchingRoutes chi.Route

	for _, r := range routes {
		if strings.Contains(r.Pattern, route) {
			matchingRoutes = r
			break
		}
	}

	return matchingRoutes
}

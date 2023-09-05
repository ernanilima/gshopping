package router_test

import (
	"fmt"
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
		getRouteByName(r.Routes(), "/v1/dashboard/total-marcas").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/dashboard/total-produtos").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/dashboard/total-produtos-nao-encontrados").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/marca").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/marca").Handlers["POST"],
		getRouteByName(r.Routes(), "/v1/marca/{id}").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/marca/{id}").Handlers["PUT"],
		getRouteByName(r.Routes(), "/v1/marca/{id}").Handlers["DELETE"],
		getRouteByName(r.Routes(), "/v1/marca/descricao/{description}").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/produto").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/produto").Handlers["POST"],
		getRouteByName(r.Routes(), "/v1/produto/{id}").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/produto/{id}").Handlers["PUT"],
		getRouteByName(r.Routes(), "/v1/produto/pesquisa/{filter}").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/produto/codigo-barras/{barcode}").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/produto/nao-encontrado").Handlers["GET"],
		getRouteByName(r.Routes(), "/v1/produto/nao-encontrado/{barcode}").Handlers["GET"],
	}

	assert.Equal(t, getTotalRoutes(r.Routes()), len(routes), "Total de rotas incompativeis")
	for index, route := range routes {
		assert.NotNil(t, route, fmt.Sprintf("deveria existir a rota de index %d", index))
	}
}

func getTotalRoutes(routes []chi.Route) int {
	var total = 0

	for _, r := range routes {
		total += len(r.Handlers)
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

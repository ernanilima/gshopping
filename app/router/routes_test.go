package router_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/ernanilima/gshopping/app/router"
// 	"github.com/go-chi/chi"
// 	"github.com/stretchr/testify/assert"
// )

// // Deve ter a rota de produto
// func TestStartRoutes_Should_Have_Product_Route(t *testing.T) {
// 	r := router.StartRoutes()

// 	// deve retornar o index da URI
// 	index := getIndexRoute(r.Routes(), "/v1/produto")

// 	assert.Greater(t, index, -1, "deve ter o index maior que -1")
// 	assert.NotNil(t, r.Routes()[index].Handlers["GET"], "deve ter o metodo GET")
// }

// // Deve retornar o status 200
// func TestStartRoutes_Should_Return_Status_200(t *testing.T) {
// 	r := router.StartRoutes()

// 	// cria uma requisicao HTTP GET para "/v1/produto"
// 	req, err := http.NewRequest("GET", "/v1/produto", nil)
// 	assert.NoError(t, err)

// 	// cria um HTTP recorder para receber a resposta
// 	res := httptest.NewRecorder()

// 	// executa a requisicao no router
// 	r.ServeHTTP(res, req)

// 	// verifica se a resposta HTTP eh 200
// 	assert.Equal(t, http.StatusOK, res.Code)
// }

// // retorna o index da rota
// func getIndexRoute(routes []chi.Route, route string) int {
// 	index := -1

// 	for i, r := range routes {
// 		if r.Pattern == route {
// 			index = i
// 			break
// 		}
// 	}

// 	return index
// }

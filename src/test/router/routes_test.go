package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernanilima/gshopping/src/app/router"
	"github.com/stretchr/testify/assert"
)

// Deve ter a rota de produto
func TestStartRoutes_Should_Have_Product_Route(t *testing.T) {
	r := router.StartRoutes()

	// deve ter a URI no index 0
	assert.Equal(t, "/v1/produto", r.Routes()[0].Pattern)
	// deve ter o metodo GET no index 0
	assert.NotNil(t, r.Routes()[0].Handlers["GET"])
}

// Deve retornar o status 200
func TestStartRoutes_Should_Return_Status_200(t *testing.T) {
	r := router.StartRoutes()

	// cria uma requisicao HTTP GET para "/v1/produto"
	req, err := http.NewRequest("GET", "/v1/produto", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	// verifica se a resposta HTTP eh 200
	assert.Equal(t, http.StatusOK, res.Code)
}

package app_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ernanilima/gshopping/app"
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/router"
	"github.com/ernanilima/gshopping/app/test/helpers"
	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	configs := helpers.GetConfigsForIntegrationTesting(context.Background())
	databaseConfig := &database.DatabaseConfig{Config: configs}
	databaseConfig.UPMigrations()

	controller := app.Init(databaseConfig)
	r := router.StartRoutes(controller)

	t.Run("Deve retornar o status 200 e um produto localizado ao pesquisar por codigo de barras", func(t *testing.T) {
		// cria uma requisicao HTTP GET para "/v1/produto/{barcode}"
		req, err := http.NewRequest("GET", "/v1/produto/789102030", nil)
		assert.NoError(t, err)
		// cria um HTTP recorder para receber a resposta
		res := httptest.NewRecorder()
		// executa a requisicao no router
		r.ServeHTTP(res, req)

		var result model.Product
		err = json.Unmarshal(res.Body.Bytes(), &result)
		assert.NoError(t, err)

		// verifica os resultados
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, result.ID)
		assert.Equal(t, "789102030", result.Barcode)
		assert.Equal(t, "Produto De Teste", result.Description)
		assert.Equal(t, "1906", result.Brand)
		assert.Equal(t, time.Now().Year(), result.CreatedAt.Year())
		assert.Equal(t, time.Now().Month(), result.CreatedAt.Month())
		assert.Equal(t, time.Now().Day(), result.CreatedAt.Day())
	})

	t.Run("Deve retornar o status 404 ao nao localizar um produto pesquisado por codigo de barras", func(t *testing.T) {
		// cria uma requisicao HTTP GET para "/v1/produto/{barcode}"
		req, err := http.NewRequest("GET", "/v1/produto/789302010", nil)
		assert.NoError(t, err)
		// cria um HTTP recorder para receber a resposta
		res := httptest.NewRecorder()
		// executa a requisicao no router
		r.ServeHTTP(res, req)

		var result response.StandardError
		err = json.Unmarshal(res.Body.Bytes(), &result)
		assert.NoError(t, err)

		// verifica os resultados
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.NotNil(t, result.Timestamp)
		assert.Equal(t, res.Code, result.Status)
		assert.Equal(t, http.StatusText(res.Code), result.Error)
		assert.Equal(t, "Produto não encontrado", result.Message)
		assert.Equal(t, "/v1/produto/789302010", result.Path)
	})
}

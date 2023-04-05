package product_controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ernanilima/gshopping/app/controller"
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/router"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var products = []model.Product{
	{
		ID:          uuid.New(),
		Barcode:     "7891020301",
		Description: "Produto para teste 1",
		Brand:       "Marda para teste 1",
		CreatedAt:   time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Barcode:     "7891020302",
		Description: "Produto para teste 2",
		Brand:       "Marda para teste 2",
		CreatedAt:   time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC),
	},
}

// Deve retornar o status 200 para buscar um produto por codigo de barras
func TestFindProductByBarcode_Should_Return_Status_200_To_Fetch_A_Product_By_Barcode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindByBarcode(gomock.Any()).Return(products[0], nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto/{barcode}"
	req, err := http.NewRequest("GET", "/v1/produto/7891020301", nil)
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
	assert.Equal(t, "7891020301", result.Barcode)
	assert.Equal(t, "Produto para teste 1", result.Description)
	assert.Equal(t, "Marda para teste 1", result.Brand)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

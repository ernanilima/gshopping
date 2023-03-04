package brand_controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernanilima/gshopping/app/controller"
	"github.com/ernanilima/gshopping/app/router"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Deve retornar o status 200 para buscar todas as marcas
func TestFindAllBrands_Should_Return_Status_200_To_Fetch_All_Brands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	repository.EXPECT().FindAll(gomock.Any()).Return(utils.Pageable{})

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto"
	req, err := http.NewRequest("GET", "/v1/marca", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	// verifica se a resposta HTTP eh 200
	assert.Equal(t, http.StatusOK, res.Code)
}

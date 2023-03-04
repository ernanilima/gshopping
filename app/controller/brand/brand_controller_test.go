package brand_controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernanilima/gshopping/app/controller"
	"github.com/ernanilima/gshopping/app/model"
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

// Deve retornar o status 200 para buscar uma marca por id
func TestFindBrandById_Should_Return_Status_200_To_Fetch_A_Brand_By_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	repository.EXPECT().FindById(gomock.Any()).Return(model.Brand{}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto"
	req, err := http.NewRequest("GET", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	// verifica se a resposta HTTP eh 200
	assert.Equal(t, http.StatusOK, res.Code)
}

// Deve retornar o status 200 para buscar todas as Marcas por descricao
func TestFindAllBrandsByDescription_Should_Return_Status_200_To_Fetch_All_Brands_By_Description(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	repository.EXPECT().FindByDescription(gomock.Any(), gomock.Any()).Return(utils.Pageable{TotalElements: 10}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto"
	req, err := http.NewRequest("GET", "/v1/marca/descricao/descri", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	// verifica se a resposta HTTP eh 200
	assert.Equal(t, http.StatusOK, res.Code)
}

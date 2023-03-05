package brand_controller_test

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
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var brands = []model.Brand{
	{
		ID:          uuid.New(),
		Description: "Marda para teste 1",
		CreatedAt:   time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Description: "Marda para teste 2",
		CreatedAt:   time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC),
	},
}

// Deve retornar o status 200 para buscar todas as marcas
func TestFindAllBrands_Should_Return_Status_200_To_Fetch_All_Brands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindAll(gomock.Any()).Return(utils.Pageable{
		Content:          brands,
		TotalPages:       0,           // total de paginas
		TotalElements:    len(brands), // total de entidades localizadas
		Size:             10,          // total de entidades por pagina
		Page:             0,           // pagina atual
		NumberOfElements: len(brands), // total de entidades por pagina
	})

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/marca"
	req, err := http.NewRequest("GET", "/v1/marca?size=12&page=0&sort=id,asc", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result utils.Pageable
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotNil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, 2, result.NumberOfElements)
}

// Deve retornar o status 200 para buscar uma marca por id
func TestFindBrandById_Should_Return_Status_200_To_Fetch_A_Brand_By_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindById(gomock.Any()).Return(brands[0], nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/marca/{id}"
	req, err := http.NewRequest("GET", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result model.Brand
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Marda para teste 1", result.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
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

package brand_controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ernanilima/gshopping/app/controller"
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/router"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/ernanilima/gshopping/app/utils/response"
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

// Deve retornar o status 422 para buscar uma marca por id quando informar o id invalido
func TestFindBrandById_Should_Return_Status_422_To_Fetch_A_Brand_By_ID_When_Informing_The_Invalid_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := controller.NewController(mocks.NewMockRepository(ctrl))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/marca/{id}"
	req, err := http.NewRequest("GET", "/v1/marca/123", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	assert.NotNil(t, result.Timestamp)
	assert.Equal(t, res.Code, result.Status)
	assert.Equal(t, http.StatusText(res.Code), result.Error)
	assert.Equal(t, "ID inválido", result.Message)
	assert.Equal(t, "/v1/marca/123", result.Path)
}

// Deve retornar o status 404 para buscar uma marca por id quando nao localizar nenhuma marca
func TestFindBrandById_Should_Return_Status_404_To_Fetch_A_Brand_By_ID_When_No_Brand_Is_Found(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindById(gomock.Any()).Return(model.Brand{}, errors.New("Error"))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/marca/{id}"
	req, err := http.NewRequest("GET", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
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
	assert.Equal(t, "Marca não encontrada", result.Message)
	assert.Equal(t, "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", result.Path)
}

// Deve retornar o status 200 para buscar marcas por descricao
func TestFindAllBrandsByDescription_Should_Return_Status_200_To_Fetch_Brands_By_Description(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindByDescription(gomock.Any(), gomock.Any()).Return(utils.Pageable{
		Content:          brands,
		TotalPages:       0,           // total de paginas
		TotalElements:    len(brands), // total de entidades localizadas
		Size:             10,          // total de entidades por pagina
		Page:             0,           // pagina atual
		NumberOfElements: len(brands), // total de entidades por pagina
	}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/marca/descricao/{description}"
	req, err := http.NewRequest("GET", "/v1/marca/descricao/Marca", nil)
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

// Deve retornar o status 404 para buscar marcas por descricao quando nao localizar nenhuma marca por erro na busca
func TestFindAllBrandsByDescription_Should_Return_Status_404_To_Fetch_Brands_By_Description_When_Not_Finding_Any_Brand_Due_To_Search_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindByDescription(gomock.Any(), gomock.Any()).Return(utils.Pageable{}, errors.New("Error"))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/marca/descricao/{description}"
	req, err := http.NewRequest("GET", "/v1/marca/descricao/nao existe", nil)
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
	assert.Equal(t, result.Status, res.Code)
	assert.Equal(t, result.Error, http.StatusText(res.Code))
	assert.Equal(t, result.Message, "Nenhuma Marca encontrada com 'nao existe'")
	assert.Equal(t, result.Path, "/v1/marca/descricao/nao existe")
}

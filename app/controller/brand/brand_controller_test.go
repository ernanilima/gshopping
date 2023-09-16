package brand_controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var brands = []model.Brand{
	{
		ID:            uuid.New(),
		Code:          11,
		Description:   "Marca para teste 1",
		TotalProducts: 5,
		CreatedAt:     time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC),
	},
	{
		ID:            uuid.New(),
		Code:          12,
		Description:   "Marca para teste 2",
		TotalProducts: 6,
		CreatedAt:     time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC),
	},
}

// Deve inserir uma marca
func TestInsertBrand_Should_Insert_A_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().InsertBrand(gomock.Any()).Return(brands[0], nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca"
	body, err := json.Marshal(model.Brand{Description: "Marca para teste 1"})
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/v1/marca", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardSuccess
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)
	dataMap, exist := result.Data.(map[string]interface{})
	assert.True(t, exist)
	jsonData, err := json.Marshal(dataMap)
	assert.NoError(t, err)
	var brand model.Brand
	err = json.Unmarshal(jsonData, &brand)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, "Marca inserida com sucesso", result.Message)
	assert.NotNil(t, brand.ID)
	assert.Equal(t, int64(11), brand.Code)
	assert.Equal(t, "Marca para teste 1", brand.Description)
	assert.Equal(t, int64(5), brand.TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brand.CreatedAt)
}

// Deve retornar o status 400 para inserir uma marca e nao enviar o body
func TestInsertBrand_Should_Return_Status_400_To_Insert_A_Brand_And_Not_Send_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca"
	req, err := http.NewRequest("POST", "/v1/marca", bytes.NewBuffer(nil))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "Erro no corpo recebido, valor inválido", result.Message)
}

// Deve retornar o status 400 para inserir uma marca com descricao muito grande
func TestInsertBrand_Should_Return_Status_400_To_Insert_A_Brand_With_A_Very_Long_Description(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca"
	body, err := json.Marshal(model.Brand{Description: "Marca para teste com descricao muito texto 1"})
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/v1/marca", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "Erro de validação: Marca inválida", result.Message)
}

// Deve retornar o status 400 para inserir uma marca que ja existe
func TestInsertBrand_Should_Return_Status_400_To_Insert_A_Brand_That_Already_Exists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().InsertBrand(gomock.Any()).Return(model.Brand{}, errors.New("Error"))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca"
	body, err := json.Marshal(model.Brand{Description: "Marca de teste 1"})
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/v1/marca", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "Marca já existe", result.Message)
}

// Deve editar uma marca
func TestEditBrand_Should_Edit_A_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().EditBrand(gomock.Any()).Return(model.Brand{
		ID:            brands[0].ID,
		Code:          brands[0].Code,
		Description:   "Marca de teste EDIT",
		TotalProducts: brands[0].TotalProducts,
		CreatedAt:     brands[0].CreatedAt,
	}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP PUT para "/v1/marca/{id}"
	body, err := json.Marshal(model.Brand{Description: "Marca de teste EDIT"})
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/v1/marca/%s", brands[0].ID), bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardSuccess
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)
	dataMap, exist := result.Data.(map[string]interface{})
	assert.True(t, exist)
	jsonData, err := json.Marshal(dataMap)
	assert.NoError(t, err)
	var brand model.Brand
	err = json.Unmarshal(jsonData, &brand)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Marca editada com sucesso", result.Message)
	assert.NotNil(t, brand.ID)
	assert.Equal(t, int64(11), brand.Code)
	assert.Equal(t, "Marca de teste EDIT", brand.Description)
	assert.Equal(t, int64(5), brand.TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brand.CreatedAt)
}

// Deve retornar o status 422 para editar uma marca e enviar o ID invalido
func TestEditBrand_Should_Return_Status_422_To_Edit_A_Brand_And_Send_Invalid_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	req, err := http.NewRequest("PUT", "/v1/marca/123", bytes.NewBuffer(nil))
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
	assert.Equal(t, "ID inválido", result.Message)
}

// Deve retornar o status 400 para editar uma marca e nao enviar o body
func TestEditBrand_Should_Return_Status_400_To_Edit_A_Brand_And_Not_Send_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	req, err := http.NewRequest("PUT", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", bytes.NewBuffer(nil))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "Erro no corpo recebido, valor inválido", result.Message)
}

// Deve retornar o status 400 para editar uma marca com descricao muito grande
func TestEditBrand_Should_Return_Status_400_To_Edit_A_Brand_With_A_Very_Long_Description(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	body, err := json.Marshal(model.Brand{Description: "Marca para teste com descricao muito texto 1"})
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "Erro de validação: Marca inválida", result.Message)
}

// Deve retornar o status 400 para editar uma marca que ja existe
func TestEditBrand_Should_Return_Status_400_To_Edit_A_Brand_That_Already_Exists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().EditBrand(gomock.Any()).Return(model.Brand{}, errors.New("Error"))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	body, err := json.Marshal(model.Brand{Description: "Marca de teste 1"})
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, "Marca já existe", result.Message)
}

// Deve deletar uma marca
func TestDeleteBrand_Should_Delete_A_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().DeleteBrand(gomock.Any()).Return(model.Brand{
		ID:            brands[0].ID,
		Code:          brands[0].Code,
		Description:   "Marca de teste DEL",
		TotalProducts: brands[0].TotalProducts,
		CreatedAt:     brands[0].CreatedAt,
	}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP DELETE para "/v1/marca/{id}"
	body, err := json.Marshal(model.Brand{Description: "Marca de teste DEL"})
	assert.NoError(t, err)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/v1/marca/%s", brands[0].ID), bytes.NewBuffer(body))
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardSuccess
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)
	dataMap, exist := result.Data.(map[string]interface{})
	assert.True(t, exist)
	jsonData, err := json.Marshal(dataMap)
	assert.NoError(t, err)
	var brand model.Brand
	err = json.Unmarshal(jsonData, &brand)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Marca excluída com sucesso", result.Message)
	assert.NotNil(t, brand.ID)
	assert.Equal(t, int64(11), brand.Code)
	assert.Equal(t, "Marca de teste DEL", brand.Description)
	assert.Equal(t, int64(5), brand.TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brand.CreatedAt)
}

// Deve retornar o status 422 para deletar uma marca e enviar o ID invalido
func TestDeleteBrand_Should_Return_Status_422_To_Delete_A_Brand_And_Send_Invalid_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	req, err := http.NewRequest("DELETE", "/v1/marca/123", bytes.NewBuffer(nil))
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
	assert.Equal(t, "ID inválido", result.Message)
}

// Deve retornar o status 404 para deletar uma marca que nao existe
func TestDeleteBrand_Should_Return_Status_404_To_Delete_A_Brand_That_Does_Not_Exist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().DeleteBrand(gomock.Any()).Return(model.Brand{}, errors.New("Error"))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	req, err := http.NewRequest("DELETE", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
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
	assert.Equal(t, "Marca não encontrada", result.Message)
}

// Deve retornar o status 409 para deletar uma marca que nao pode ser removida
func TestDeleteBrand_Should_Return_Status_409_To_Delete_A_Brand_That_Cannot_Be_Removed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().DeleteBrand(gomock.Any()).Return(model.Brand{}, &pq.Error{Message: "Error", Code: "unique_violation"})

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/marca/{id}"
	req, err := http.NewRequest("DELETE", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusConflict, res.Code)
	assert.Equal(t, "Marca não pode ser removida", result.Message)
}

// Deve retornar o status 200 para buscar todas as marcas
func TestFindAllBrands_Should_Return_Status_200_To_Fetch_All_Brands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindAllBrands(gomock.Any()).Return(utils.Pageable{
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
	repository.EXPECT().FindBrandById(gomock.Any()).Return(brands[0], nil)

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
	assert.Equal(t, "Marca para teste 1", result.Description)
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
	repository.EXPECT().FindBrandById(gomock.Any()).Return(model.Brand{}, errors.New("Error"))

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
	repository.EXPECT().FindAllBrandsByDescription(gomock.Any(), gomock.Any()).Return(utils.Pageable{
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
	repository.EXPECT().FindAllBrandsByDescription(gomock.Any(), gomock.Any()).Return(utils.Pageable{}, errors.New("Error"))

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

// Deve retornar o status 404 para buscar marcas por descricao quando nao localizar nenhuma marca por erro na busca
func TestFindAllBrandsByDescription_Should_Return_Status_404_To_Fetch_Brands_By_Description_When_Not_Finding_Any_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindAllBrandsByDescription(gomock.Any(), gomock.Any()).Return(utils.Pageable{}, nil)

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

// Deve retornar o total de marcas
func TestFindTotalBrands_Should_Return_Total_Brands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindTotalBrands().Return(int32(15))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/dashboard/total-marcas"
	req, err := http.NewRequest("GET", "/v1/dashboard/total-marcas", nil)
	assert.NoError(t, err)

	// cria um HTTP recorder para receber a resposta
	res := httptest.NewRecorder()

	// executa a requisicao no router
	r.ServeHTTP(res, req)

	var result int32
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, int32(15), result)
}

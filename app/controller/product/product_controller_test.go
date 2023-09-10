package product_controller_test

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
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var products = []model.Product{
	{
		ID:          uuid.New(),
		Barcode:     "7891020301",
		Description: "Produto para teste 1",
		Brand:       model.Brand{ID: uuid.New(), Description: "Marca para teste 1"},
		CreatedAt:   time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Barcode:     "7891020302",
		Description: "Produto para teste 2",
		Brand:       model.Brand{ID: uuid.New(), Description: "Marca para teste 2"},
		CreatedAt:   time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC),
	},
}

var productsNotFound = []model.ProductNotFound{
	{
		ID:       1,
		Barcode:  "229874455661",
		Attempts: 111,
	},
	{
		ID:       2,
		Barcode:  "229874455662",
		Attempts: 12,
	},
}

// Deve inserir um produto
func TestInsertProduct_Should_Insert_A_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().InsertProduct(gomock.Any()).Return(products[0], nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP POST para "/v1/produto"
	body, err := json.Marshal(model.Product{
		Barcode:     "7891020301",
		Description: "Produto para teste 1",
		Brand:       model.Brand{ID: uuid.New()},
	})
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/v1/produto", bytes.NewBuffer(body))
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
	var product model.Product
	err = json.Unmarshal(jsonData, &product)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, "Produto inserido com sucesso", result.Message)
	assert.NotNil(t, product.ID)
	assert.Equal(t, "7891020301", product.Barcode)
	assert.Equal(t, "Produto para teste 1", product.Description)
	assert.NotNil(t, product.Brand.ID)
	assert.Equal(t, "Marca para teste 1", product.Brand.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), product.CreatedAt)
}

// Deve editar um produto
func TestEditProduct_Should_Edit_A_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().EditProduct(gomock.Any()).Return(model.Product{
		ID:          products[0].ID,
		Barcode:     "7890000001",
		Description: "Produto de teste EDIT",
		Brand:       model.Brand{ID: uuid.New(), Description: "Marca de teste EDIT"},
		CreatedAt:   products[0].CreatedAt,
	}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP PUT para "/v1/produto/{id}"
	body, err := json.Marshal(model.Product{
		ID:          products[0].ID,
		Barcode:     "7890000001",
		Description: "Produto de teste EDIT",
		Brand:       model.Brand{ID: uuid.New(), Description: "Marca de teste EDIT"},
		CreatedAt:   products[0].CreatedAt,
	})
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/v1/produto/%s", products[0].ID), bytes.NewBuffer(body))
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
	var product model.Product
	err = json.Unmarshal(jsonData, &product)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Produto editado com sucesso", result.Message)
	assert.NotNil(t, product.ID)
	assert.Equal(t, "7890000001", product.Barcode)
	assert.Equal(t, "Produto de teste EDIT", product.Description)
	assert.NotNil(t, product.Brand.ID)
	assert.Equal(t, "Marca de teste EDIT", product.Brand.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), product.CreatedAt)
}

// Deve retornar o status 200 para buscar todos os produtos
func TestFindAllProducts_Should_Return_Status_200_To_Fetch_All_Products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindAllProducts(gomock.Any(), gomock.Any()).Return(utils.Pageable{
		Content:          products,
		TotalPages:       0,             // total de paginas
		TotalElements:    len(products), // total de entidades localizadas
		Size:             10,            // total de entidades por pagina
		Page:             0,             // pagina atual
		NumberOfElements: len(products), // total de entidades por pagina
	})

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto"
	req, err := http.NewRequest("GET", "/v1/produto?size=12&page=0&sort=id,asc", nil)
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

// Deve retornar o status 200 para buscar um produto por id
func TestFindProductById_Should_Return_Status_200_To_Fetch_A_Product_By_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindProductById(gomock.Any()).Return(products[0], nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto/{id}"
	req, err := http.NewRequest("GET", "/v1/produto/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
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
	assert.NotNil(t, result.Brand.ID)
	assert.Equal(t, "Marca para teste 1", result.Brand.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

// Deve retornar o status 200 para buscar um produto por codigo de barras
func TestFindProductByBarcode_Should_Return_Status_200_To_Fetch_A_Product_By_Barcode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindProductByBarcode(gomock.Any()).Return(products[0], nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto/{barcode}"
	req, err := http.NewRequest("GET", "/v1/produto/codigo-barras/7891020301", nil)
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
	assert.NotNil(t, result.Brand.ID)
	assert.Equal(t, "Marca para teste 1", result.Brand.Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result.CreatedAt)
}

// Deve retornar o status 404 para buscar um produto por codigo de barras quando nao localizar nenhum produto
func TestFindProductByBarcode_Should_Return_Status_404_To_Fetch_A_Product_By_Barcode_When_No_Product_Is_Found(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindProductByBarcode(gomock.Any()).Return(model.Product{}, errors.New("Error"))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto/{barcode}"
	req, err := http.NewRequest("GET", "/v1/produto/codigo-barras/7891020301", nil)
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
	assert.Equal(t, "/v1/produto/codigo-barras/7891020301", result.Path)
}

// Deve retornar o status 200 para buscar todos os produtos nao encontrados
func TestFindAllProductsNotFound_Should_Return_Status_200_To_Fetch_All_ProductsNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindAllProductsNotFound(gomock.Any()).Return(utils.Pageable{
		Content:          productsNotFound,
		TotalPages:       0,                     // total de paginas
		TotalElements:    len(productsNotFound), // total de entidades localizadas
		Size:             10,                    // total de entidades por pagina
		Page:             0,                     // pagina atual
		NumberOfElements: len(productsNotFound), // total de entidades por pagina
	})

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto/nao-encontrado"
	req, err := http.NewRequest("GET", "/v1/produto/nao-encontrado?size=12&page=0&sort=id,asc", nil)
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

// Deve retornar o status 200 para buscar todos os produtos nao encontrados por codigo de barras
func TestFindAllProductsNotFoundByBarcode_Should_Return_Status_200_To_Fetch_All_ProductsNotFound_By_Barcode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindAllProductsNotFoundByBarcode(gomock.Any(), gomock.Any()).Return(utils.Pageable{
		Content:          productsNotFound,
		TotalPages:       0,                     // total de paginas
		TotalElements:    len(productsNotFound), // total de entidades localizadas
		Size:             10,                    // total de entidades por pagina
		Page:             0,                     // pagina atual
		NumberOfElements: len(productsNotFound), // total de entidades por pagina
	}, nil)

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/produto/nao-encontrado/{barcode}"
	req, err := http.NewRequest("GET", "/v1/produto/nao-encontrado/229874455", nil)
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

// Deve retornar o total de produtos
func TestFindTotalProducts_Should_Return_Total_Products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindTotalProducts().Return(int32(16))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/dashboard/total-produtos"
	req, err := http.NewRequest("GET", "/v1/dashboard/total-produtos", nil)
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
	assert.Equal(t, int32(16), result)
}

// Deve retornar o total de produtos cadastrados como nao encontrados
func TestFindTotalProductsNotFound_Should_Return_Total_ProductsNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockRepository(ctrl)
	controller := controller.NewController(repository)
	repository.EXPECT().FindTotalProductsNotFound().Return(int32(17))

	r := router.StartRoutes(controller)

	// cria uma requisicao HTTP GET para "/v1/dashboard/total-produtos-nao-encontrados"
	req, err := http.NewRequest("GET", "/v1/dashboard/total-produtos-nao-encontrados", nil)
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
	assert.Equal(t, int32(17), result)
}

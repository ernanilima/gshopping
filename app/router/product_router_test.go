package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernanilima/gshopping/app/router"
	"github.com/ernanilima/gshopping/app/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Deve acessar a rota e retornar status code 200 ao inserir um produto
func TestRouteInsertProduct_Should_Access_The_Route_And_Return_Status_Code_200_When_Insert_A_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("POST", "/v1/produto", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), req.RequestURI)
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao editar um produto
func TestRouteEditProduct_Should_Access_The_Route_And_Return_Status_Code_200_When_Edit_A_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().EditProduct(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("PUT", "/v1/produto/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), "/v1/produto/{id}")
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar todos os produtos
func TestRouteFindAllProducts_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_All_Products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindAllProducts(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), req.RequestURI)
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar todos os produtos pro filtro
func TestRouteFindAllProducts_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_All_Product_By_Filter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindAllProducts(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto/pesquisa/789111222333", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), "/v1/produto/pesquisa/{filter}")
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar um produto por id
func TestRouteFindProductById_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_A_Product_By_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindProductById(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), "/v1/produto/{id}")
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar produto por codigo de barras
func TestRouteFindProductByBarcode_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_A_Product_By_Barcode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindProductByBarcode(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto/codigo-barras/789102030", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), "/v1/produto/codigo-barras/{barcode}")
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar todos os produtos nao encontrados
func TestRouteFindAllProductsNotFound_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_All_ProductsNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindAllProductsNotFound(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto/nao-encontrado", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), req.RequestURI)
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar todos os produtos nao encontrados por codigo de barras
func TestRouteFindAllProductsNotFoundByBarcode_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_All_ProductsNotFound_By_Barcode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindAllProductsNotFoundByBarcode(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto/nao-encontrado/789102030", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), "/v1/produto/nao-encontrado/{barcode}")
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

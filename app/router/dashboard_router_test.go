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

// Deve acessar a rota e retornar status code 200 ao buscar o total de marcas
func TestRouteFindTotalBrands_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_Total_Brands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindTotalBrands(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/dashboard/total-marcas", nil)
	res := httptest.NewRecorder()

	routeDashboardMarca := getRouteByName(r.Routes(), req.RequestURI)
	routeDashboardMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar o total de produtos
func TestRouteFindTotalProducts_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_Total_Products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindTotalProducts(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/dashboard/total-produtos", nil)
	res := httptest.NewRecorder()

	routeDashboardProduto := getRouteByName(r.Routes(), req.RequestURI)
	routeDashboardProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar o total de produtos nao encontrados
func TestRouteFindTotalProductsNotFound_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_Total_ProductsNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindTotalProductsNotFound(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/dashboard/total-produtos-nao-encontrados", nil)
	res := httptest.NewRecorder()

	routeDashboardProdutoNaoEncontrado := getRouteByName(r.Routes(), req.RequestURI)
	routeDashboardProdutoNaoEncontrado.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

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

// Deve acessar a rota e retornar status code 200 ao inserir uma marca
func TestRouteInsertBrand_Should_Access_The_Route_And_Return_Status_Code_200_When_Insert_A_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().InsertBrand(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("POST", "/v1/marca", nil)
	res := httptest.NewRecorder()

	routeMarca := getRouteByName(r.Routes(), req.RequestURI)
	routeMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao editar uma marca
func TestRouteEditBrand_Should_Access_The_Route_And_Return_Status_Code_200_When_Edit_A_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().EditBrand(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("PUT", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	res := httptest.NewRecorder()

	routeMarca := getRouteByName(r.Routes(), "/v1/marca/{id}")
	routeMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao deletar uma marca
func TestRouteDeleteBrand_Should_Access_The_Route_And_Return_Status_Code_200_When_Delete_A_Brand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().DeleteBrand(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("DELETE", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	res := httptest.NewRecorder()

	routeMarca := getRouteByName(r.Routes(), "/v1/marca/{id}")
	routeMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar todas as marcas
func TestRouteFindAllBrands_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_All_Brands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindAllBrands(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/marca", nil)
	res := httptest.NewRecorder()

	routeMarca := getRouteByName(r.Routes(), req.RequestURI)
	routeMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar uma marca por id
func TestRouteFindBrandById_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_A_Brand_By_ID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindBrandById(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/marca/6f75b5bc-e561-4bc7-a28d-e74bc706a4e9", nil)
	res := httptest.NewRecorder()

	routeMarca := getRouteByName(r.Routes(), "/v1/marca/{id}")
	routeMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

// Deve acessar a rota e retornar status code 200 ao buscar todas as marcas por descricao
func TestRouteFindAllBrandsByDescription_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_All_Brands_By_Description(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindAllBrandsByDescription(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/marca/descricao/descri", nil)
	res := httptest.NewRecorder()

	routeMarca := getRouteByName(r.Routes(), "/v1/marca/descricao/{description}")
	routeMarca.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

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

// Deve acessar a rota e retornar status code 200 ao buscar produto por codigo de barras
func TestRouteFindProductByBarcode_Should_Access_The_Route_And_Return_Status_Code_200_When_Fetch_A_Product_By_Barcode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	controller := mocks.NewMockController(ctrl)
	controller.EXPECT().FindProductByBarcode(gomock.Any(), gomock.Any()).AnyTimes().Return()

	r := router.StartRoutes(controller)

	req := httptest.NewRequest("GET", "/v1/produto/789102030", nil)
	res := httptest.NewRecorder()

	routeProduto := getRouteByName(r.Routes(), "/v1/produto/{barcode}")
	routeProduto.Handlers[req.Method].ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "deveria retornar o status code 200")
}

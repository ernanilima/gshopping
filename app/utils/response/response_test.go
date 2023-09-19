package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var brands = []model.Brand{
	{
		ID:            uuid.New(),
		Code:          11,
		Description:   "Marca para teste 1",
		TotalProducts: 6,
		CreatedAt:     time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC),
	},
	{
		ID:            uuid.New(),
		Code:          12,
		Description:   "Marca para teste 2",
		TotalProducts: 7,
		CreatedAt:     time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC),
	},
}

// Deve retornar um request padrao
func TestJSON_Should_Return_A_Default_Request(t *testing.T) {
	res := httptest.NewRecorder()

	response.JSON(res, http.StatusOK, brands)

	var result []model.Brand
	err := json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, 1, len(res.Header()))
	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Equal(t, 2, len(result))
	assert.NotNil(t, result[0].ID)
	assert.Equal(t, int64(11), result[0].Code)
	assert.Equal(t, "Marca para teste 1", result[0].Description)
	assert.Equal(t, int64(6), result[0].TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result[0].CreatedAt)
	assert.NotNil(t, result[1].ID)
	assert.Equal(t, int64(12), result[1].Code)
	assert.Equal(t, "Marca para teste 2", result[1].Description)
	assert.Equal(t, int64(7), result[1].TotalProducts)
	assert.Equal(t, time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC), result[1].CreatedAt)
}

// Deve retornar um request usado para sucesso
func TestSuccess_Should_Return_A_Request_For_Success(t *testing.T) {
	res := httptest.NewRecorder()
	messageSuccess := "Marca inserida com sucesso"

	response.Success(res, http.StatusCreated, brands[0], messageSuccess)

	var result response.StandardSuccess
	err := json.Unmarshal(res.Body.Bytes(), &result)
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
	assert.Equal(t, int64(6), brand.TotalProducts)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), brand.CreatedAt)
}

// Deve retornar um request usado para erro
func TestError_Should_Return_A_Request_For_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/marca?size=12&page=0&sort=id,asc", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	messageError := "Não encontrada"

	response.Error(res, req, http.StatusNotFound, messageError)

	var result response.StandardError
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.NoError(t, err)

	// verifica os resultados
	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.NotNil(t, result.Timestamp)
	assert.Equal(t, res.Code, result.Status)
	assert.Equal(t, http.StatusText(res.Code), result.Error)
	assert.Equal(t, "Não encontrada", result.Message)
	assert.Equal(t, "/v1/marca", result.Path)
}

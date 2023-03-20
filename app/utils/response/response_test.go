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

// Deve retornar um request usado para sucesso
func TestJSON_Should_Return_A_Request_For_Success(t *testing.T) {
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
	assert.Equal(t, "Marda para teste 1", result[0].Description)
	assert.Equal(t, time.Date(2021, time.January, 1, 21, 31, 41, 0, time.UTC), result[0].CreatedAt)
	assert.NotNil(t, result[1].ID)
	assert.Equal(t, "Marda para teste 2", result[1].Description)
	assert.Equal(t, time.Date(2022, time.February, 2, 22, 32, 42, 0, time.UTC), result[1].CreatedAt)
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

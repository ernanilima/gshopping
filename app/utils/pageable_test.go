package utils_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/ernanilima/gshopping/app/utils"
	"github.com/stretchr/testify/assert"
)

// Deve retornar os dados de paginacao padrao quando nao informar filtros de paginacao
func TestPaginationFilters_Should_Return_Default_Pagination_Data_When_Not_Providing_Pagination_Filters(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

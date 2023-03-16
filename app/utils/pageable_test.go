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

// Deve retornar os dados de paginacao padrao com alteracao apenas em page quando informar apenas esse parametro
func TestPaginationFilters_Should_Return_Default_Pagination_Data_With_Change_Only_In_PAGE_When_Informing_Only_This_Parameter(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "page=2",
		},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 10, result.Size)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

// Deve retornar os dados de paginacao padrao com alteracao apenas em size quando informar apenas esse parametro
func TestPaginationFilters_Should_Return_Default_Pagination_Data_With_Change_Only_In_SIZE_When_Informing_Only_This_Parameter(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "size=20",
		},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

// Deve retornar os dados de paginacao padrao sem alteracao em sort quando o parameto for incompleto
func TestPaginationFilters_Should_Return_Default_Pagination_Data_Without_Changing_SORT_When_Parameter_Is_Incomplete(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "size=20&page=2&sort=description", // apenas um valor de ordenacao (nesse caso foi o campo)
		},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

// Deve retornar os dados de paginacao padrao sem alteracao em sort quando o parameto for para ordenacao invalida
func TestPaginationFilters_Should_Return_Default_Pagination_Data_Without_Changing_SORT_When_Parameter_Is_For_Invalid_Sort(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "size=20&page=2&sort=description,crescente", // aceita apenas 'asc' ou 'desc'
		},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, "id ASC", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

package utils_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var brands = []model.Brand{
	{
		ID:          uuid.New(),
		Description: "Marda para teste 1",
		CreatedAt:   time.Date(2020, time.January, 1, 21, 31, 41, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Description: "Marda para teste 2",
		CreatedAt:   time.Date(2020, time.January, 2, 22, 32, 42, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Description: "Marda para teste 3",
		CreatedAt:   time.Date(2020, time.January, 3, 23, 33, 43, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Description: "Marda para teste 4",
		CreatedAt:   time.Date(2020, time.January, 4, 24, 34, 44, 0, time.UTC),
	},
	{
		ID:          uuid.New(),
		Description: "Marda para teste 5",
		CreatedAt:   time.Date(2020, time.January, 5, 25, 35, 45, 0, time.UTC),
	},
}

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

// Deve retornar os dados de paginacao conforme os dados passados por parametro (sort asc)
func TestPaginationFilters_Should_Return_Pagination_Data_According_To_The_Data_Passed_By_Parameter_SORT_ASC(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "size=20&page=2&sort=description,asc",
		},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, "description asc", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

// Deve retornar os dados de paginacao conforme os dados passados por parametro (sort desc)
func TestPaginationFilters_Should_Return_Pagination_Data_According_To_The_Data_Passed_By_Parameter_SORT_DESC(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "size=20&page=2&sort=description,desc",
		},
	}

	result := utils.PaginationFilters(request)

	// verifica os resultados
	assert.Nil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 0, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, "description desc", result.Sort)
	assert.Equal(t, 0, result.NumberOfElements)
}

// Deve retornar um Pageable com os dados que devem ser exibidos em um request para uma entidade
func TestGeneratePaginationRequest_Should_Return_A_Pageable_With_The_Data_That_Must_Be_Displayed_In_A_Request_For_An_Entity(t *testing.T) {

	content := []model.Brand{brands[0]}

	pageable := utils.Pageable{
		TotalElements: len(content), // total de entidades localizadas
		Size:          20,           // total de entidades por pagina
		Page:          0,            // pagina atual
	}

	result := utils.GeneratePaginationRequest(content, pageable)

	// verifica os resultados
	assert.NotNil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 1, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, 1, result.NumberOfElements)
}

// Deve retornar um Pageable com os dados que devem ser exibidos em um request para diversas entidades na mesma pagina
func TestGeneratePaginationRequest_Should_Return_A_Pageable_With_The_Data_That_Must_Be_Displayed_In_A_Request_For_Several_Entities_On_The_Same_Page(t *testing.T) {

	content := brands

	pageable := utils.Pageable{
		TotalElements: 5,  // total de entidades localizadas (no sql)
		Size:          20, // total de entidades por pagina
		Page:          0,  // pagina atual
	}

	result := utils.GeneratePaginationRequest(content, pageable)

	// verifica os resultados
	assert.NotNil(t, result.Content)
	assert.Equal(t, 0, result.TotalPages)
	assert.Equal(t, 5, result.TotalElements)
	assert.Equal(t, 20, result.Size)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, 5, result.NumberOfElements) // total enditades que serao exibidas
}

// Deve retornar um Pageable com os dados que devem ser exibidos em um request para diversas entidades para diversas paginas
func TestGeneratePaginationRequest_Should_Return_A_Pageable_With_The_Data_That_Must_Be_Displayed_In_A_Request_For_Several_Entities_For_Several_Pages(t *testing.T) {

	content := brands

	pageable := utils.Pageable{
		TotalElements: 11, // total de entidades localizadas
		Size:          3,  // total de entidades por pagina
		Page:          1,  // pagina atual
	}

	result := utils.GeneratePaginationRequest(content, pageable)

	// verifica os resultados
	assert.NotNil(t, result.Content)
	assert.Equal(t, 3, result.TotalPages)
	assert.Equal(t, 11, result.TotalElements)
	assert.Equal(t, 3, result.Size)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 5, result.NumberOfElements) // total enditades que serao exibidas
}

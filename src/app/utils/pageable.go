package utils

import (
	"net/http"
	"reflect"
	"strconv"
)

var (
	defaultSize = 10
	defaultPage = 0
)

// Pageable eh o objeto de retorna para requisicoes com paginacao
type Pageable struct {
	Content          interface{} `json:"content"`
	TotalPages       int         `json:"totalPages"`
	TotalElements    int         `json:"totalElements"`
	Size             int         `json:"size"`
	Page             int         `json:"page"`
	NumberOfElements int         `json:"numberOfElements"`
}

// PaginationFilters monta o objeto de paginacao para
// realizar a busca no banco de dados
func PaginationFilters(r *http.Request) Pageable {
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = defaultSize
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = defaultPage
	}

	return Pageable{
		Size: size,
		Page: page,
	}
}

// GeneratePaginationRequest monta o objeto Pageable para retorno
func GeneratePaginationRequest(content interface{}, pageable Pageable) Pageable {
	pageable.Content = content
	pageable.TotalPages = pageable.TotalElements / pageable.Size
	pageable.NumberOfElements = reflect.ValueOf(content).Len()
	return pageable
}

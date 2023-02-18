package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var (
	defaultSize = 10
	defaultPage = 0
	defaultSort = "id ASC"
)

// Pageable eh o objeto de retorna para requisicoes com paginacao
type Pageable struct {
	Content          interface{} `json:"content"`
	TotalPages       int         `json:"totalPages"`
	TotalElements    int         `json:"totalElements"`
	Size             int         `json:"size"`
	Page             int         `json:"page"`
	Sort             string      `json:"-"`
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
		Sort: sortParameter(r),
	}
}

// sortParameter monta a ordenacao
func sortParameter(r *http.Request) string {
	sortParam := r.URL.Query().Get("sort")
	if sortParam == "" {
		return defaultSort
	}

	params := strings.Split(sortParam, ",")
	if len(params) != 2 {
		return defaultSort
	}

	if strings.ToLower(params[1]) != "asc" && strings.ToLower(params[1]) != "desc" {
		return defaultSort
	}

	return fmt.Sprintf("%s %s", params[0], params[1])
}

// GeneratePaginationRequest monta o objeto Pageable para retorno
func GeneratePaginationRequest(content interface{}, pageable Pageable) Pageable {
	pageable.Content = content
	pageable.TotalPages = pageable.TotalElements / pageable.Size
	pageable.NumberOfElements = reflect.ValueOf(content).Len()
	return pageable
}

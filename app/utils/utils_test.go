package utils_test

import (
	"testing"

	"github.com/ernanilima/gshopping/app/utils"
	"github.com/stretchr/testify/assert"
)

// Deve retornar o diretorio principal da aplicacao para testes
func TestGetURIPath_Should_Return_Main_Directory_Of_The_Application_For_Testing(t *testing.T) {
	path := utils.GetURIPath()

	// verifica os resultados
	assert.NotNil(t, path)
	assert.Contains(t, path, "/gshopping/")
}

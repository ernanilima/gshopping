package utils_test

import (
	"os"
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

// Deve retornar o diretorio raiz quando Getwd retornar barra
func TestGetURIPath_Should_Return_Root_Directory_When_Getwd_Returns_Slash(t *testing.T) {
	originalDir, err := os.Getwd()
	assert.NoError(t, err)
	defer func() { os.Chdir(originalDir) }()

	// Muda o diretorio de trabalho para "/"
	os.Chdir("/")

	path := utils.GetURIPath()

	// verifica os resultados
	assert.NotNil(t, path)
	assert.Equal(t, "/", path)
}

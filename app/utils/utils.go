package utils

import (
	"os"
	"strings"
)

// GetURIPath retorna o diretorio da aplicacao
func GetURIPath() string {
	rootDir, _ := os.Getwd()
	if rootDir == "/" {
		// quando executar a aplicacao normalmente
		return rootDir
	}
	// quando executar em testes
	index := strings.LastIndex(rootDir, "/gshopping")
	return rootDir[:index+len("/gshopping/")]
}

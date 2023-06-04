package utils

import (
	"os"
	"strings"
)

// GetURIPath retorna o diretorio da aplicacao
func GetURIPath() string {
	rootDir, _ := os.Getwd()
	if rootDir == "/" {
		// quando executa a aplicacao pelo docker-compose
		return rootDir
	}

	index := strings.LastIndex(rootDir, "/gshopping")

	if strings.Contains(rootDir, "src/gshopping") {
		// quando executar a aplicacao normalmente
		return rootDir[:index+len("/gshopping")] + "/"
	}

	// quando executar em testes
	return rootDir[:index+len("/gshopping/")]
}

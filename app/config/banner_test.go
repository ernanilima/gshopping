package config_test

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/stretchr/testify/assert"
)

var configs = config.Config{
	Server: struct {
		Port    int    `mapstructure:"port"`
		Version string `mapstructure:"version"`
	}{
		Port:    8989,
		Version: "1.0",
	},
}

// Deve imprimir um erro por nao encontrar o banner
func TestStartBanner_Should_Print_A_Error_For_Not_Finding_The_Banner(t *testing.T) {
	// capturar a saida do metodo
	output := CaptureOutput(func() {
		config.StartBanner(configs)
	})

	assert.Contains(t, output, "Erro ao ler arquivo banner.txt")
}

// Deve imprimir o banner
func TestStartBanner_Should_Print_The_Banner(t *testing.T) {
	bannerForTest("CREATE")

	// capturar a saida do metodo
	output := CaptureOutput(func() {
		config.StartBanner(configs)
	})

	bannerForTest("DELETE")

	expected := fmt.Sprintf("Banner - :: Go: %s :: Porta: %d (v%s)\n",
		runtime.Version(),
		configs.Server.Port,
		configs.Server.Version)

	assert.Equal(t, expected, output)
}

func bannerForTest(action string) {
	filename := "banner.txt"
	if action == "CREATE" {
		content := []byte("Banner - :: Go: %s :: Porta: %d (v%s)")
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Erro ao criar o arquivo:", err)
			return
		}
		defer file.Close()
		if _, err = file.Write(content); err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}
	} else if action == "DELETE" {
		if os.Remove(filename) != nil {
			fmt.Println("Erro ao apagar o arquivo:")
			return
		}
	}
}

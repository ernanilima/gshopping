package config

import (
	"fmt"
	"os"
	"runtime"
)

// StartBanner exibe o banner com alguns dados da aplicacao
func StartBanner(configs Config) {
	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo banner.txt: ", err)
	}

	fmt.Printf("%s\n", fmt.Sprintf(string(banner),
		runtime.Version(),
		configs.Server.Port,
		configs.Server.Version))
}

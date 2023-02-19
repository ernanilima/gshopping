package config

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

// StartBanner exibe o banner com alguns dados da aplicacao
func StartBanner(configs Config) {
	banner, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo banner.txt: ", err)
	}

	fmt.Printf("%s\n", fmt.Sprintf(string(banner),
		runtime.Version(),
		configs.Server.Port,
		configs.Server.Version))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/ernanilima/gshopping/src/app"
	"github.com/ernanilima/gshopping/src/app/config"
	"github.com/ernanilima/gshopping/src/app/router"
)

func main() {
	configs := config.GetConfigs()
	routes := router.StartRoutes()

	displayAPIConnection(configs)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Server.Port), routes))
}

// displayAPIConnection exibe onde acessar a API
func displayAPIConnection(configs config.Config) {

	horizontalLine := strings.Repeat("-", 35)
	message := fmt.Sprintf("%s\n Aplicao iniciando na porta: %d\n%s\n", horizontalLine, configs.Server.Port, horizontalLine)

	fmt.Printf(message)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/ernanilima/gshopping/src/app"
	"github.com/ernanilima/gshopping/src/app/config"
	"github.com/ernanilima/gshopping/src/app/router"
)

func main() {
	configs := config.GetConfigs()
	routes := router.StartRoutes()

	config.StartBanner(configs)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Server.Port), routes))
}

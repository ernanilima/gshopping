package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/ernanilima/gshopping/app"
	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/router"
)

func main() {
	configs := config.GetConfigs()
	routes := router.StartRoutes()

	config.StartBanner(configs)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Server.Port), routes))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/ernanilima/gshopping/src/app/config"
	. "github.com/ernanilima/gshopping/src/app/router"
	"github.com/go-chi/chi"
)

var routes *chi.Mux

func init() {
	StartConfig()

	routes = StartRoutes()
}

func main() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", CONFIG.Server.Port), routes))
}

package main

import (
	"log"
	"net/http"

	"github.com/ernanilima/gshopping/src/app/router"
)

func main() {
	r := router.StartRoutes()

	log.Fatal(http.ListenAndServe(":4000", r))
}

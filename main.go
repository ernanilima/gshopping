package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ernanilima/gshopping/app"
	_ "github.com/ernanilima/gshopping/app"
	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/router"
)

func main() {
	// carrega as configuracoes da aplicacao
	var config = &config.Config{}
	config = config.StartConfig(".")

	// realiza uma conexao com o banco de dados
	conn, err := database.OpenConnection(config)
	if err != nil {
		log.Fatalf("falha ao conectar com o banco de dados: %s", err)
	}

	// executa as migrations pendentes
	database.UPMigrations(conn)
	conn.Close()

	controller := app.Init(config)
	routes := router.StartRoutes(controller)

	config.StartBanner()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), routes))
}

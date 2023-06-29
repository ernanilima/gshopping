package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ernanilima/gshopping/app"
	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/router"
	"github.com/rs/cors"
)

func main() {
	// carrega as configuracoes da aplicacao
	var configs = &config.Config{}
	configs = configs.StartConfig(".")

	// captura as configuracoes do banco de dados
	databaseConfig := &database.DatabaseConfig{Config: configs}
	// executa todas as migracoes
	databaseConfig.UPMigrations()

	// inicializa os repositories e os controllers
	controller := app.Init(databaseConfig)
	// inicia as rotas
	routes := router.StartRoutes(controller)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler(routes)

	// inicia o banner
	configs.StartBanner()

	// executa a aplicacao
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Server.Port), handler))
}

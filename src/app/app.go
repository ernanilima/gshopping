package app

import (
	"log"

	"github.com/ernanilima/gshopping/src/app/config"
	"github.com/ernanilima/gshopping/src/app/repository"
)

func init() {
	// carrega as configuracoes da aplicacao
	config.StartConfig()

	// realiza uma conexao com o banco de dados
	conn, err := repository.OpenConnection()
	if err != nil {
		log.Fatalf("falha ao conectar com o banco de dados: %s", err)
	}

	// executa as migrations pendentes
	repository.UPMigrations(conn)
	conn.Close()
}

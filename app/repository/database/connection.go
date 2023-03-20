package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ernanilima/gshopping/app/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

// OpenConnection abre uma conecao com o banco de dados
func OpenConnection(configs *config.Config) *sql.DB {
	configPostgres := configs.Database.Postgres

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configPostgres.Host,
		configPostgres.Port,
		configPostgres.User,
		configPostgres.Pass,
		configPostgres.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("falha ao conectar com o banco de dados: %s", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("falha ao conectar com o banco de dados: %s", err)
	}

	return db
}

// UPMigrations executa as migrations pendentes
func UPMigrations(db *sql.DB) {
	if err := goose.Up(db, "./db/postgres"); err != nil {
		log.Fatalf("falha ao aplicar migrations para postgres: %s", err)
	}

	if err := goose.Up(db, "./db/migrations"); err != nil {
		log.Fatalf("falha ao aplicar migrations: %s", err)
	}
}

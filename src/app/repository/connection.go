package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ernanilima/gshopping/src/app/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

// OpenConnection abre uma conecao com o banco de dados
func OpenConnection() (*sql.DB, error) {
	configs := config.GetConfigs()
	configPostgres := configs.Database.Postgres

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configPostgres.Host,
		configPostgres.Port,
		configPostgres.User,
		configPostgres.Pass,
		configPostgres.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("falha ao conectar com o postgres: %s", err)
		return nil, err
	}

	return db, db.Ping()
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

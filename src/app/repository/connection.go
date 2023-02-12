package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ernanilima/gshopping/src/app/config"
	_ "github.com/lib/pq"
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

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	defer db.Close()

	return db, err
}

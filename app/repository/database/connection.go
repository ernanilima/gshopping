package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/utils"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type DatabaseConnector interface {
	OpenConnection() *sql.DB
	UPMigrations()
}

type DatabaseConfig struct {
	*config.Config
}

// OpenConnection abre uma conecao com o banco de dados
func (databaseConfig *DatabaseConfig) OpenConnection() *sql.DB {
	configPostgres := databaseConfig.Database.Postgres

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
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
func (databaseConfig *DatabaseConfig) UPMigrations() {
	db := databaseConfig.OpenConnection()
	defer db.Close()

	if err := goose.Up(db, utils.GetURIPath()+"db/migrations"); err != nil {
		log.Fatalf("falha ao aplicar migrations: %s", err)
	}
}

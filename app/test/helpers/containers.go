package helpers

import (
	"context"
	"log"
	"time"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	username = "postgres"
	password = "postgres"
	database = "gshopping-test"
)

// GetConfigsForIntegrationTesting cria um banco de dados postgres para teste de integracao com testcontainers
// https://golang.testcontainers.org/modules/postgres
// https://github.com/testcontainers/tc-guide-getting-started-with-testcontainers-for-go/blob/main/customer/repo_test.go
// https://medium.com/@dilshataliev/integration-tests-with-golang-test-containers-and-postgres-abb49e8096c5
func GetConfigsForIntegrationTesting(ctx context.Context) *config.Config {
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.1-alpine"),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
		postgres.WithDatabase(database),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal("Erro ao criar container: ", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatal("Erro ao obter host do container: ", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal("Erro ao obter port do container: ", err)
	}

	configs := new(config.Config)
	configs.Database.Postgres = struct {
		Host string "mapstructure:\"db_host\""
		Port string "mapstructure:\"db_port\""
		User string "mapstructure:\"db_user\""
		Pass string "mapstructure:\"db_pass\""
		Name string "mapstructure:\"db_name\""
	}{
		Host: host,
		Port: port.Port(),
		User: username,
		Pass: password,
		Name: database,
	}
	return configs
}

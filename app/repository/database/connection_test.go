package database_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/test/helpers"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestOpenConnection_Success(t *testing.T) {
	configs := helpers.GetConfigsForIntegrationTesting(context.Background(), t)

	databaseConfig := &database.DatabaseConfig{Config: configs}
	conn := databaseConfig.OpenConnection()
	defer conn.Close()

	// verifica os resultados
	assert.NotNil(t, conn)
}

func TestUPMigrations_Success(t *testing.T) {
	configs := helpers.GetConfigsForIntegrationTesting(context.Background(), t)

	databaseConfig := &database.DatabaseConfig{Config: configs}

	output := CaptureOutput(func() {
		databaseConfig.UPMigrations()
	})

	// verifica os resultados
	assert.NotContains(t, output, "falha ao aplicar migrations")
}

// Funcao para capturar a saida da funcao passada por parametro
func CaptureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf := bytes.Buffer{}
	io.Copy(&buf, r)
	return buf.String()
}

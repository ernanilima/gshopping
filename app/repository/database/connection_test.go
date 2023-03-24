package database_test

import (
	"context"
	"testing"

	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/test/helpers"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestOpenConnection_Success(t *testing.T) {
	configs := helpers.GetConfigsForIntegrationTesting(context.Background())

	conn := database.OpenConnection(configs)
	defer conn.Close()

	// verifica os resultados
	assert.NotNil(t, conn)
}

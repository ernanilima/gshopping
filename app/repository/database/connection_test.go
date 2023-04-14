package database_test

import (
	_ "github.com/lib/pq"
)

// func TestOpenConnection_Success(t *testing.T) {
// 	configs := helpers.GetConfigsForIntegrationTesting(context.Background())

// 	databaseConfig := &database.DatabaseConfig{Config: configs}
// 	conn := databaseConfig.OpenConnection()
// 	defer conn.Close()

// 	// verifica os resultados
// 	assert.NotNil(t, conn)
// }

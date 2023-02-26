package config_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/stretchr/testify/assert"
)

var configYML = []byte(`
server:
  version: "0.1"
  port: 4040
database:
  postgres:
    db_host: ${DB_HOST}
    db_port: ${DB_PORT}
    db_user: ${DB_USER}
    db_pass: ${DB_PASS}
    db_name: ${DB_NAME}
`)

var dotENV = []byte(`
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres_u
DB_PASS=postgres_p
DB_NAME=db
`)

// Deve testar as funcoes do metodo
func TestStartConfig(t *testing.T) {
	// Deve retornar configuracoes baseadas em dados de variaveis de ambiente do sistema operacional
	t.Run("TestStartConfig_Should_Return_Settings_Based_On_Data_From_Operating_System_Environment_Variables", func(t *testing.T) {
		tempDir := t.TempDir()
		var configs = &config.Config{}

		// defina variaveis de ambiente temporarias para o sistema operacional
		os.Setenv("DB_HOST", "localhost_os")
		os.Setenv("DB_PORT", "54320")
		os.Setenv("DB_USER", "postgres_u_os")
		os.Setenv("DB_PASS", "postgres_p_os")
		os.Setenv("DB_NAME", "db_os")
		defer os.Unsetenv("DB_HOST")
		defer os.Unsetenv("DB_PORT")
		defer os.Unsetenv("DB_USER")
		defer os.Unsetenv("DB_PASS")
		defer os.Unsetenv("DB_NAME")

		// crie um arquivo config.yml temporario
		os.WriteFile(fmt.Sprintf("%s/config.yml", tempDir), configYML, 0777)

		// capturar a saida do metodo
		output := CaptureOutput(func() {
			configs = configs.StartConfig(tempDir)
		})

		assert.Equal(t, output, "Variaveis de ambiente serao carregadas do sistema operacional\n")

		assert.Equal(t, 4040, configs.Server.Port)
		assert.Equal(t, "0.1", configs.Server.Version)
		assert.Equal(t, "localhost_os", configs.Database.Postgres.Host)
		assert.Equal(t, "54320", configs.Database.Postgres.Port)
		assert.Equal(t, "postgres_u_os", configs.Database.Postgres.User)
		assert.Equal(t, "postgres_p_os", configs.Database.Postgres.Pass)
		assert.Equal(t, "db_os", configs.Database.Postgres.Name)
	})

	// Deve retornar configuracoes baseadas em dados de variaveis de ambiente do .env
	t.Run("TestStartConfig_Should_Return_Settings_Based_On_ENV_Environment_Variables_Data", func(t *testing.T) {
		tempDir := t.TempDir()
		var configs = &config.Config{}

		// crie um arquivo .env temporario
		tmpDotEnvFile := createTempFile(".env", dotENV)
		defer os.Remove(tmpDotEnvFile)
		// crie um arquivo config.yml temporario
		os.WriteFile(fmt.Sprintf("%s/config.yml", tempDir), configYML, 0777)

		// capturar a saida do metodo
		output := CaptureOutput(func() {
			configs = configs.StartConfig(tempDir)
		})

		assert.Equal(t, output, "")

		assert.Equal(t, 4040, configs.Server.Port)
		assert.Equal(t, "0.1", configs.Server.Version)
		assert.Equal(t, "localhost", configs.Database.Postgres.Host)
		assert.Equal(t, "5432", configs.Database.Postgres.Port)
		assert.Equal(t, "postgres_u", configs.Database.Postgres.User)
		assert.Equal(t, "postgres_p", configs.Database.Postgres.Pass)
		assert.Equal(t, "db", configs.Database.Postgres.Name)
	})

	// Deve retornar um erro por nao localizar o arquivo config.yml
	t.Run("TestStartConfig_Should_Return_Error_For_Not_Finding_The_ConfigYML_File", func(t *testing.T) {
		assert.Panics(t, func() { configs.StartConfig(t.TempDir()) })
	})
}

// Funcao para criar um arquivo temporario
func createTempFile(filename string, content []byte) string {
	tmpfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	if _, err := tmpfile.Write(content); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}
	return tmpfile.Name()
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

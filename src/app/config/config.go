package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var CONFIG *config

// estrutura das configuracoes
type config struct {
	Server struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
		Port    int    `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		URI  string `mapstructure:"db_uri"`
		User string `mapstructure:"db_user"`
		Pass string `mapstructure:"db_pass"`
	} `mapstructure:"database"`
}

// StartConfig inicia a construcao das configuracoes
func StartConfig() {
	viper.AddConfigPath(getPwd())
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar as variaveis de ambiente: %s", err)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Erro ao ler as configuracoes: %s", err)
	}

	CONFIG = new(config)
	if err := viper.Unmarshal(&CONFIG); err != nil {
		log.Fatalf("Nao foi possivel decodificar o arquivo de configuracao: %s", err)
	}

	CONFIG.Database = struct {
		URI  string "mapstructure:\"db_uri\""
		User string "mapstructure:\"db_user\""
		Pass string "mapstructure:\"db_pass\""
	}{
		URI:  os.ExpandEnv(CONFIG.Database.URI),
		User: os.ExpandEnv(CONFIG.Database.User),
		Pass: os.ExpandEnv(CONFIG.Database.Pass),
	}
}

// Retorna o diretorio do projeto
func getPwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter o diretorio do projeto: %s", err)
	}
	return dir
}

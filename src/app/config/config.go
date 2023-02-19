package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *Config

// estrutura das configuracoes
type Config struct {
	Server struct {
		Port    int    `mapstructure:"port"`
		Version string `mapstructure:"version"`
	} `mapstructure:"server"`
	Database struct {
		Postgres struct {
			Host string `mapstructure:"db_host"`
			Port string `mapstructure:"db_port"`
			User string `mapstructure:"db_user"`
			Pass string `mapstructure:"db_pass"`
			Name string `mapstructure:"db_name"`
		} `mapstructure:"postgres"`
	} `mapstructure:"database"`
}

// StartConfig inicia a construcao das configuracoes
func StartConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := godotenv.Load(); err != nil {
		fmt.Println("Variaveis de ambiente serao carregadas do sistema operacional")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler as configuracoes: %s", err)
	}

	cfg = new(Config)
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Nao foi possivel decodificar o arquivo de configuracao: %s", err)
	}

	cfg.Database.Postgres = struct {
		Host string "mapstructure:\"db_host\""
		Port string "mapstructure:\"db_port\""
		User string "mapstructure:\"db_user\""
		Pass string "mapstructure:\"db_pass\""
		Name string "mapstructure:\"db_name\""
	}{
		Host: os.ExpandEnv(cfg.Database.Postgres.Host),
		Port: os.ExpandEnv(cfg.Database.Postgres.Port),
		User: os.ExpandEnv(cfg.Database.Postgres.User),
		Pass: os.ExpandEnv(cfg.Database.Postgres.Pass),
		Name: os.ExpandEnv(cfg.Database.Postgres.Name),
	}
}

// GetConfigs retorna as configuracoes da aplicacao
func GetConfigs() Config {
	return *cfg
}

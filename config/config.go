package config

import (
	"os"
)

// Config armazena as variáveis de ambiente da aplicação
type Config struct {
	APPENV     string
	APPPort    string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSync     bool
}

// NewConfig retorna uma nova instância de Config
func NewConfig() *Config {
	return &Config{
		APPENV:     os.Getenv("ENV"),
		APPPort:    os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBSync:     os.Getenv("DB_SYNC") == "true",
	}
}

package main

import (
	"github.com/georgejr3211/food-flow/config"
	"github.com/georgejr3211/food-flow/internal"
	"github.com/georgejr3211/food-flow/internal/utils/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error.Fatalf("Erro ao carregar as variáveis de ambiente: %v", err)
	}

	logger.Info.Println("Starting Food Flow API...", config.NewConfig().APPENV)

	app := internal.NewApp()
	server := internal.NewServer(app)
	server.Start()
}

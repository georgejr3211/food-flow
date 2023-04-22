package main

import (
	"github.com/georgejr3211/food-flow/config"
	_ "github.com/georgejr3211/food-flow/docs"
	"github.com/georgejr3211/food-flow/internal/infrastructure/http"
	"github.com/georgejr3211/food-flow/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Food Flow API
// @description API para o projeto Food Flow
// @version 1
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error.Fatalf("Erro ao carregar as variáveis de ambiente: %v", err)
	}

	logger.Info.Println("Starting Food Flow API...", config.NewConfig().APPENV)

	router := gin.Default()

	app := http.NewApp(router)
	app.Init()
	app.Run()
}

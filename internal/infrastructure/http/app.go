package http

import (
	"github.com/georgejr3211/food-flow/config"
	gormconfig "github.com/georgejr3211/food-flow/internal/infrastructure/database/gorm/config"
	"github.com/georgejr3211/food-flow/internal/utils/logger"
	"github.com/gin-gonic/gin"
)

// App é a estrutura principal da aplicação
type App struct {
	router *gin.Engine
}

// NewApp retorna uma nova instância de App
func NewApp(router *gin.Engine) *App {
	return &App{
		router: router,
	}
}

// Init inicia a aplicação
func (a *App) Init() {

	db := gormconfig.ConnectDB()
	logger.Info.Println("Connected to database")

	router := NewRouter(a.router)
	router.Init(db)

	// Define os middlewares da aplicação
	// a.Router.Use(LoggerMiddleware)
}

// Run executa a aplicação
func (a *App) Run() {
	config := config.NewConfig()
	APPPort := config.APPPort
	APPHost := config.APPHost + ":" + APPPort
	a.router.Run(APPHost)
}

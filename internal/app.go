package internal

import (
	"net/http"

	"github.com/georgejr3211/food-flow/config"
	gormconfig "github.com/georgejr3211/food-flow/internal/infrastructure/database/gorm/config"
	"github.com/georgejr3211/food-flow/internal/module/product"
	"github.com/georgejr3211/food-flow/internal/utils/logger"
	"github.com/gorilla/mux"
)

// App é a estrutura principal da aplicação
type App struct {
	Router *mux.Router
}

// NewApp retorna uma nova instância de App
func NewApp() *App {
	return &App{}
}

// Init inicia a aplicação
func (a *App) Init() {
	a.Router = mux.NewRouter()

	db := gormconfig.ConnectDB()
	logger.Info.Println("Connected to database")

	productRoute := product.InitProductModule(db)
	productRoute.RegisterRoutes(a.Router)

	// Define os middlewares da aplicação
	// a.Router.Use(LoggerMiddleware)
}

// Run executa a aplicação
func (a *App) Run() {
	APPPort := config.NewConfig().APPPort
	APPHost := "0.0.0.0" + ":" + APPPort
	logger.Info.Printf("Starting server on port %v\n", APPPort)
	http.ListenAndServe(APPHost, a.Router)
}

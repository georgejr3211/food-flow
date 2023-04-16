package internal

import (
	"github.com/georgejr3211/food-flow/internal/utils/logger"
)

// Server é a estrutura que armazena o servidor
type Server struct {
	App *App
}

// NewServer retorna uma nova instância de Server
func NewServer(app *App) *Server {
	return &Server{App: app}
}

// Start inicia o servidor
func (s *Server) Start() {
	// Inicializa a aplicação
	s.App.Init()

	// Executa a aplicação
	s.App.Run()
}

// Stop finaliza o servidor
func (s *Server) Stop() {
	logger.Info.Println("Stopping server...")
	// TODO: Implementar lógica de finalização do servidor
}

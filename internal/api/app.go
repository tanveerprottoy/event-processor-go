package api

import (
	"os"

	"github.com/tanveerprottoy/event-processor-go/pkg/server"
)

// App contains the components of the application
type App struct {
	cfg *Config
	srv *server.Server
}

// NewApp creates App
func NewApp() *App {
	a := &App{cfg: NewConfig()}
	
	a.initServer()
	a.configureGracefulShutdown()

	return a
}

// initServer initializes the server
func (a *App) initServer() {
	a.srv = server.NewServer(":"+os.Getenv("PORT"), a.cfg.router.Mux)
}

// configureGracefulShutdown configures graceful shutdown
func (a *App) configureGracefulShutdown() {
	a.srv.ConfigureGracefulShutdown(func() {
		
	})
}

// Start starts the server
func (a *App) Start() {
	a.srv.Start()
}

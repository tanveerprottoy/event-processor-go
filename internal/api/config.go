package api

import (
	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/route"
	filecfg "github.com/tanveerprottoy/event-processor-go/internal/api/file/config"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"github.com/tanveerprottoy/event-processor-go/pkg/env"
	"github.com/tanveerprottoy/event-processor-go/pkg/router"
)

type Config struct {
	router *router.Router
}

func NewConfig() *Config {
	c := new(Config)
	c.loadEnv()
	c.initRouter()
	c.initComponents()
	return c
}

// loadEnv initializes env
func (c *Config) loadEnv() {
	env.LoadEnv("")
}

// initRouter initializes router
func (c *Config) initRouter() {
	c.router = router.NewRouter()
}

func (c *Config) initRoutes(args ...any) {
	// static routes
	route.Static(c.router.Mux, constant.ApiPattern+"/v1/static", args[0].(*handler.Static))

	//file routes
	route.File(c.router.Mux, constant.ApiPattern+"/v1/files", args[1].(*handler.File))
}

// initComponents initializes application components
func (c *Config) initComponents() {
	file := filecfg.NewConfig()

	c.initRoutes(
		handler.NewStatic(),
		handler.NewFile(file.UseCase),
	)
}

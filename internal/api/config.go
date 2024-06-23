package api

import (
	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/route"
	filecfg "github.com/tanveerprottoy/event-processor-go/internal/api/file/config"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"github.com/tanveerprottoy/event-processor-go/pkg/router"
	"github.com/tanveerprottoy/event-processor-go/pkg/env"
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
	route.File(constant.ApiPattern+"/v1", args[0].(*handler.File))
}

// initComponents initializes application components
func (c *Config) initComponents() {
	file := filecfg.NewConfig()
	c.initRoutes(
		handler.NewFile(file.UseCase),
	)
}

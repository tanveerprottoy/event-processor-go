package filecfg

import (
	"github.com/tanveerprottoy/event-processor-go/internal/api/file"
	"github.com/tanveerprottoy/event-processor-go/internal/api/file/service"
)

// Config holds the components of the current package
type Config struct {
	UseCase file.UseCase
}

// NewConfig initializes a new NewConfig
func NewConfig() *Config {
	u := service.NewService()
	return &Config{UseCase: u}
}

package route

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
)

func File(basePattern string, handler *handler.File) {
	http.HandleFunc("POST "+basePattern+"/upload", handler.Upload)
}

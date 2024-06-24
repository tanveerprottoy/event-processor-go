package route

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
)

func File(basePattern string, handler *handler.File) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.HandleFunc("POST "+basePattern+"/upload", handler.Upload)
	})
}

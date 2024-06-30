package route

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
)

func Static(mux *http.ServeMux, basePattern string, handler *handler.Static) {
	mux.HandleFunc("GET "+basePattern+"/form", handler.ServeHTML)
}

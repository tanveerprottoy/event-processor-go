package route

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
)

func File(mux *http.ServeMux, basePattern string, handler *handler.File) {
	mux.HandleFunc("POST "+basePattern+"/upload", handler.Upload)
	mux.HandleFunc("POST "+basePattern+"/upload-multi", handler.UploadMultiple)
}

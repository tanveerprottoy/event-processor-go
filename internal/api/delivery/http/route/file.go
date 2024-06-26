package route

import (
	"net/http"

	"github.com/tanveerprottoy/event-processor-go/internal/api/delivery/http/handler"
)

func File(handler *handler.File) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.HandleFunc("/files/upload", handler.Upload)
	})
}

package router

import (
	"net/http"
)

// Router struct
type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	r := &Router{}
	r.Mux = http.NewServeMux()
	// r.registerGlobalMiddlewares()
	// r.registerNotFoundHander()
	return r
}

/* func (r *Router) registerGlobalMiddlewares() {
	r.Mux.

		middleware.Logger,
		middleware.Recoverer,
		middlewarext.JSONContentTypeMiddleWare,
		// timeout middlewares
		middleware.Timeout(constant.RequestTimeout*time.Second),
		middlewarext.TimeoutHandler(constant.RequestTimeout*time.Second),
		middlewarext.CORSEnableMiddleWare,
	)
}

func (r *Router) registerNotFoundHander() {
	// not found handler
	r.Mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\": \"Resource not found\"}"))
	})
}
*/

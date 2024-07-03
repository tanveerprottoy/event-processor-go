package httpext

import (
	"net/http"
)

func GetURLParam(r *http.Request, key string) string {
	return r.PathValue(key)
}

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

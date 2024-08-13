package middleware

import "net/http"

// Middleware is a type that wraps an http.HandlerFunc
// it takes in one handlerfunc and wraps it within another handlerfunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain builds the middlware chain recursively
func Chain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](Chain(f, m[1:cap(m)]...))
}

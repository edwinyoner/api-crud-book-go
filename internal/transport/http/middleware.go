package http

import "net/http"

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("➡️ ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

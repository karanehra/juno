package middlewares

import (
	"fmt"
	"net/http"
)

//LoggerMiddleware outputs request logs for all incoming request
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s : %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

//JSONContentMiddleware adds a content type header to all outgoing responses
func JSONContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

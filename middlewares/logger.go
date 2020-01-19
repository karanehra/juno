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

//CORSMiddleware adds the access control headers to requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		res.Header().Add("Access-Control-Allow-Origin", "http://localhost:8000")
		res.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, PUT")
		if req.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(res, req)
	})
}

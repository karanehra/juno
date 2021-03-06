package middlewares

import (
	"juno/util"
	"net/http"
	"strings"
)

//AuthMiddleware extracts a token from incoming requests and authenticates it
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		if !strings.Contains(uri, "user") {
			header := r.Header.Get("Authorization")
			if header == "" {
				util.SendUnauthorizedResponse(w, "Auth not provided")
				return
			}
			token := strings.Split(header, " ")[1]
			if !util.IsJWTValid(token) {
				util.SendUnauthorizedResponse(w, "Invalid JWT")
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

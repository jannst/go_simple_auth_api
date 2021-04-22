package middlewares

import (
	"net/http"
)

func SimpleCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Allow", "OPTIONS, POST, GET, PUT")
			w.Header().Set("Access-Control-Allow-Headers", "X-API-TOKEN, Content-Type")
			w.WriteHeader(http.StatusNoContent)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
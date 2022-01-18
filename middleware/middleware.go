package middleware

import "net/http"

func SetContentType(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		inner.ServeHTTP(w, r)
	})
}

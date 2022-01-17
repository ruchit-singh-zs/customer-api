package main

import (
	"github.com/anushi/customer-api/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customers/{id}", handler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customers", handler.Post).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.Use(SetContentType)

	log.Fatal(srv.ListenAndServe())
}

func SetContentType(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		inner.ServeHTTP(w, r)
	})
}

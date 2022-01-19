package main

import (
	"customer-api/handler"
	"customer-api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customers/{id}", handler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customers", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/delete/{id}", handler.DeleteByID).Methods(http.MethodDelete)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.Use(middleware.SetContentType)

	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"customer-api/handlers"
	"customer-api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customer/{id}", handlers.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customer", handlers.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/delete/{id}", handlers.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/customer/update/{id}", handlers.UpdateByID).Methods(http.MethodPut)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.Use(middleware.SetContentType)

	log.Fatal(srv.ListenAndServe())
}

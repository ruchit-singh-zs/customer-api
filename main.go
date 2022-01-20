package main

import (
	"customer-api/handlers"
	"customer-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customer/{id}", handlers.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customer", handlers.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/delete/{id}", handlers.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/customer/update/{id}", handlers.UpdateByID).Methods(http.MethodPut)

	r.Use(middleware.SetContentType)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Cant Connect!")
	}
}

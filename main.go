package main

import (
	"customer-api/handlers/customer"
	"customer-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
// 	r.HandleFunc("/customer/{id}", customer.GetByID).Methods(http.MethodGet)
// 	r.HandleFunc("/customer", customer.Create).Methods(http.MethodPost)
// 	r.HandleFunc("/customer/delete/{id}", customer.DeleteByID).Methods(http.MethodDelete)
// 	r.HandleFunc("/customer/update/{id}", customer.UpdateByID).Methods(http.MethodPut)

	
	r.Use(middleware.SetContentType)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Cant Connect!")
	}
}

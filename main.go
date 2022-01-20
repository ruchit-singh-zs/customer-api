package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"customer-api/drivers"
	"customer-api/handlers"
	"customer-api/middleware"
	"customer-api/stores"
)

func main() {

	db, _ := drivers.ConnectToSQL()
	store := stores.New(db)
	h := handlers.New(store)
	r := mux.NewRouter()

	r.HandleFunc("/customer/{id}", h.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customer/create", handlers.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/delete/{id}", handlers.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/customer/update/{id}", handlers.UpdateByID).Methods(http.MethodPut)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Cant connect to localhost")
		return
	}

	r.Use(middleware.SetContentType)
}

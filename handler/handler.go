package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"customer-api/driver"
	"customer-api/model"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	db := driver.ConnectToSQL()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var customer model.Customer

	err := db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.PhoneNo, &customer.Address)

	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
	case nil:
		resp, err := json.Marshal(customer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		_, err = w.Write(resp)
		if err != nil {
			log.Println(err)
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var c model.Customer
	db := driver.ConnectToSQL()
	defer db.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err2 := json.Unmarshal(body, &c)

	if err2 != nil {
		return
	}

	_, err = db.Exec("INSERT INTO Customer (ID,NAME , PHONENO, ADDRESS) VALUES (?,?, ?, ?)",
		c.ID, c.Name, c.PhoneNo, c.Address)

	if err != nil {
		log.Printf("Error in Inserting: %v", err)
	}
	log.Println(c)
}
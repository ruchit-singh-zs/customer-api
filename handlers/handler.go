package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"customer-api/drivers"
	"customer-api/models"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	db := drivers.ConnectToSQL()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var customer models.Customer

	err := db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.PhoneNo, &customer.Address)

	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte("No Record Exists"))
		if err != nil {
			return
		}
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

func Create(w http.ResponseWriter, r *http.Request) {
	var c models.Customer
	db := drivers.ConnectToSQL()
	defer db.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	error := json.Unmarshal(body, &c)
	if error != nil {
		return
	}

	_, err = db.Exec("INSERT INTO Customer (ID,NAME , PHONENO, ADDRESS) VALUES (?,?, ?, ?)",
		&c.ID, &c.Name, &c.PhoneNo, &c.Address)

	if err != nil {
		log.Printf("Error in Inserting: %v", err)
	}
	log.Println(c)
	_, err = w.Write([]byte("Succesfully created"))
	if err != nil {
		return
	}
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := drivers.ConnectToSQL()
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM Customer WHERE ID =?", id)
	if err != nil {
		log.Println("Error in deleting", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write([]byte("Deleted Successfully"))
	if err != nil {
		return
	}
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	var c models.Customer
	db := drivers.ConnectToSQL()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	error := json.Unmarshal(body, &c)
	if error != nil {
		return
	}

	_, err = db.Exec("UPDATE Customer SET NAME = ?, PHONENO=?, ADDRESS=? WHERE ID = ?",
		&c.Name, &c.PhoneNo, &c.Address, id)

	if err != nil {
		log.Printf("Error in Updating: %v", err)
	}

	_, err = w.Write([]byte("Updated Successfully"))
	if err != nil {
		return
	}
}

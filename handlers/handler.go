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
	"customer-api/stores"
)

type Handler struct {
	store stores.Customer
}

func New(s stores.Customer) Handler {
	return Handler{store: s}
}

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var customer models.Customer
	err = db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.PhoneNo, &customer.Address)

	//customer, err := h.store.GetCustomer(id)
	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Record Exists"))
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
	db, err := drivers.ConnectToSQL()
	if err != nil {
		log.Println("Connection Failed")
	}
	defer db.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var c models.Customer
	error := json.Unmarshal(body, &c)
	if error != nil {
		log.Println("Cannot encode the data")
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = db.Exec("INSERT INTO Customer (ID,NAME , PHONENO, ADDRESS) VALUES (?,?, ?, ?)",
		&c.ID, &c.Name, &c.PhoneNo, &c.Address)

	if err != nil {
		log.Printf("Error in Inserting: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println(c)
	_, err = w.Write([]byte("Succesfully created"))
	if err != nil {
		log.Println("HTTP Reply not working")
	}
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		log.Println("Connection lost")
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	_, err = db.Exec("DELETE FROM Customer WHERE ID =?", id)
	if err != nil {
		log.Println("Error in deleting", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte("Deleted Successfully"))
	if err != nil {
		//w.WriteHeader(http.StatusNoContent)
		return
	}
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		log.Println("Connection lost")
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var c models.Customer
	error := json.Unmarshal(body, &c)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE Customer SET NAME = ?, PHONENO=?, ADDRESS=? WHERE ID = ?",
		&c.Name, &c.PhoneNo, &c.Address, id)

	if err != nil {
		log.Printf("Error in Updating: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write([]byte("Updated Successfully"))
	if err != nil {
		return
	}
}

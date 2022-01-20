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

func Create(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var c models.Customer
	err = json.Unmarshal(body, &c)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO Customer (ID,NAME , PHONENO, ADDRESS) VALUES (?,?, ?, ?)",
		&c.ID, &c.Name, &c.PhoneNo, &c.Address)

	if err != nil {
		log.Printf("Error in Inserting: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	_, err = w.Write([]byte("successfully created"))
	if err != nil {
		log.Println(err)
	}
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var c models.Customer

	err = db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&c.ID, &c.Name, &c.PhoneNo, &c.Address)

	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte("No Record Exists"))

		if err != nil {
			log.Println(err)
		}
	case nil:
		resp, err := json.Marshal(c)
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

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var c models.Customer
	err = json.Unmarshal(body, &c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE Customer SET NAME = ?, PHONENO=?, ADDRESS=? WHERE ID = ?",
		&c.Name, &c.PhoneNo, &c.Address, id)

	if err != nil {
		log.Printf("Error in Updating: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = w.Write([]byte("Updated Successfully"))
	if err != nil {
		log.Println(err)
	}
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	_, err = db.Exec("DELETE FROM Customer WHERE ID =?", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in deleting", err)

		return
	}

	_, err = w.Write([]byte("Deleted Successfully"))
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}
}

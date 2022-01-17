package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Customer struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	PhoneNo string `json:"phoneNo"`
	Address string `json:"address"`
}

var db *sql.DB

func setUpDB() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "organisation",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db

}
func main() {
	db = setUpDB()
	r := mux.NewRouter()
	r.HandleFunc("/retrieve/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/post", Post).Methods(http.MethodPost)

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

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var customer Customer

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

		w.WriteHeader(http.StatusOK)
		_, err1 := w.Write(resp)
		if err1 != nil {
			return
		}

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func Post(w http.ResponseWriter, r *http.Request) {
	var c Customer
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

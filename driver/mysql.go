package driver

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToSQL() *sql.DB {
	// capture connection properties
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "organization",
	}

	// get a database handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected!")

	return db
}

package drivers

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToSQL() (*sql.DB, error) {
	// capture connection properties
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "organisation",
	}

	// get a database handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	log.Println("Connected!")

	return db, nil
}

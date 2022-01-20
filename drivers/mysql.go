package drivers

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
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
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected!")

	return db, nil
}

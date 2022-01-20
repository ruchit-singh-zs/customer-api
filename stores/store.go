package stores

import (
	"customer-api/errors"
	"customer-api/models"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCustomer(id string) (models.Customer, error) {

	var customer models.Customer

	err := s.db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.PhoneNo, &customer.Address)

	switch err {
	case sql.ErrNoRows:
		return models.Customer{}, errors.EntityNotFound{Entity: "customer", ID: id}
	case nil:
		return customer, nil
	default:
		return models.Customer{}, errors.DB{Err: err}
	}
}

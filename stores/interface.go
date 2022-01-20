package stores

import "customer-api/models"

type Customer interface {
	GetCustomer(id string) (models.Customer, error)
}

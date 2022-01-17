package model

type Customer struct {
	ID      int64  `json:"customer_id"`
	Name    string `json:"name"`
	PhoneNo string `json:"phone_no"`
	Address string `json:"address"`
}

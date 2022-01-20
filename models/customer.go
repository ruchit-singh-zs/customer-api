package models

type Customer struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	PhoneNo string `json:"phoneNo"`
	Address string `json:"address"`
}

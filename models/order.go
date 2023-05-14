package models

type Order struct {
	ID          int `json:"id"`
	Fecha       int `json:"fecha"`
	Customer_Id int `json:"customer_id"`
}

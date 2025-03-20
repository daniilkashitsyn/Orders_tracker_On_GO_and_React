package models

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Adress      string `json:"adress"`
	Ration      string `json:"ration"`
	PhoneNumber string `json:"phone_number"`
}

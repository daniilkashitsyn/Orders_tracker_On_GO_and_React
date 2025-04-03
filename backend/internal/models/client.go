package models

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Ration      string `json:"ration"`
	PhoneNumber string `json:"phone_number"`
}

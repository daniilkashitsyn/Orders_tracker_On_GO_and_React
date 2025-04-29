package database

import (
	"backend/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() error {
	var err error
	DB, err = sql.Open("sqlite3", "../database/base.db")

	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Successfully connected to database")

	return nil
}

func GetClients(query string) ([]models.Client, error) {
	var clients []models.Client

	rows, err := DB.Query(query)

	if err != nil {
		fmt.Printf("error querying clients: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("error closing rows: %v", err)
		}
	}(rows)
	for rows.Next() {
		var client models.Client

		err := rows.Scan(&client.ID, &client.Name, &client.Address, &client.Ration, &client.PhoneNumber)

		if err != nil {
			fmt.Printf("error scanning row: %v", err)
			continue
		}
		var clientAdd models.Client

		clientAdd.ID = client.ID
		clientAdd.Name = client.Name
		clientAdd.Address = client.Address
		clientAdd.Ration = client.Ration
		clientAdd.PhoneNumber = client.PhoneNumber

		clients = append(clients, clientAdd)
	}
	return clients, nil
}

func GetClientByID(id string) (models.Client, error) {
	var client models.Client

	err := DB.QueryRow("select * from clients where id=$1", id).Scan(&client.ID, &client.Name, &client.Address,
		&client.Ration, &client.PhoneNumber)

	return client, err
}

func DeleteClient(id string) error {
	_, err := DB.Exec("delete from clients where id=?", id)

	if err != nil {
		fmt.Printf("error deleting client: %v", err)
		return err
	}

	return nil
}

func CreateClient(client models.Client) error {
	_, err := DB.Exec("insert into clients(name, address, ration, phone_number) values(?, ?, ?, ?)",
		client.Name, client.Address, client.Ration, client.PhoneNumber,
	)

	if err != nil {
		fmt.Printf("error inserting client: %v", err)
		return err
	}

	return nil
}

func UpdateClient(client models.Client) error {

	_, err := DB.Exec("update clients set name=?, address=?, ration=?, phone_number=? where id=?",
		client.Name, client.Address, client.Ration, client.PhoneNumber, client.ID,
	)

	if err != nil {
		fmt.Printf("error updating client: %v", err)
		return err
	}

	return nil
}

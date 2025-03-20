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

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	rows, err := DB.Query("select * from books")

	if err != nil {
		fmt.Printf("error querying books: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)

		if err != nil {
			fmt.Printf("error scanning row: %v", err)
			continue
		}
		var book models.Book

		book.ID = id
		book.Title = name

		books = append(books, book)
	}
	return books, nil
}

func AddBook(book models.Book) error {
	_, err := DB.Exec("insert into books(id, title) values(?, ?)", book.ID, book.Title)

	if err != nil {
		return fmt.Errorf("error inserting book: %v", err)
	}

	return nil
}

func GetClients(query string) ([]models.Client, error) {
	var clients []models.Client

	rows, err := DB.Query(query)
	// rows, err := DB.Query("select * from clients order by ration")

	if err != nil {
		fmt.Printf("error querying clients: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var client models.Client

		err := rows.Scan(&client.ID, &client.Name, &client.Adress, &client.Ration, &client.PhoneNumber)

		if err != nil {
			fmt.Printf("error scanning row: %v", err)
			continue
		}
		var clientAdd models.Client

		clientAdd.ID = client.ID
		clientAdd.Name = client.Name
		clientAdd.Adress = client.Adress
		clientAdd.Ration = client.Ration
		clientAdd.PhoneNumber = client.PhoneNumber

		clients = append(clients, clientAdd)
	}
	return clients, nil
}

package database

import (
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

func GetBooks() {
	rows, err := DB.Query("select * from books")

	if err != nil {
		fmt.Printf("error querying books: %v", err)
		return
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
		fmt.Printf("id: %v, name: %v\n", id, name)
	}

}

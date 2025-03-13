package main

import (
	"backend/internal/database"
	"fmt"
)

func main() {
	err := database.InitDb()
	if err != nil {
		fmt.Println(err)
	}

	database.GetBooks()
}

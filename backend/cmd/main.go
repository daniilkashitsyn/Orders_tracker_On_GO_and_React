package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @title React + Go Orders tracker

func main() {
	err := database.InitDb()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	router.Group("/*")
	{
		router.GET(
			"/clients",
			handlers.GetClients,
		)

		router.DELETE(
			"/client/del/:id",
			handlers.DeleteClient,
		)

		router.POST(
			"/clients/add",
			handlers.CreateClient,
		)
	}

	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

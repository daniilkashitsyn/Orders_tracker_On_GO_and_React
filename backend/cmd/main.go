package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/auth"
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

	router.POST(
		"/login",
		auth.Login,
	)
	protected := router.Group("/")
	//router.Use(func(c *gin.Context) {
	//	c.Header("Access-Control-Allow-Origin", "*")
	//})
	protected.Use(auth.AuthMiddleware())
	protected.Use(auth.RoleBasedAccess())

	protected.Group("/*")
	{
		protected.GET(
			"/clients",
			handlers.GetClients,
		)

		protected.DELETE(
			"/clients/del/:id",
			handlers.DeleteClient,
		)

		protected.POST(
			"/clients/add",
			handlers.CreateClient,
		)

		protected.PUT(
			"/clients/upd/:id",
			handlers.UpdateClient,
		)

		protected.GET(
			"/clients/:id",
			handlers.GetClientById,
		)
	}

	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

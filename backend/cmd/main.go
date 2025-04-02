package main

import (
	_ "backend/docs"
	"backend/internal/database"
	"backend/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		//router.DELETE(
		//	"/clients",
		//	handlers.DeleteClient,
		//)

		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

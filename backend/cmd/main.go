package main

import (
	"backend/internal/database"
	"fmt"
	"net/http"

	_ "backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @titile React + Go Orders tracker

// GetBooks godoc
// @Summary Получить список книг
// @Description Возвращает все книги из базы данных
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string
// @Router /api/books [get]
func main() {
	err := database.InitDb()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	router.Group("/*")
	{
		router.GET("/books", getBooks)
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Run(":8080")
}

func getBooks(c *gin.Context) {
	books, err := database.GetBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, books)
}

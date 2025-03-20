package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBooks godoc
// @Summary Получить список книг
// @Description Возвращает все книги из базы данных
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string
// @Router /books [get]
func GetBooks(c *gin.Context) {
	books, err := database.GetBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, books)
}

// AddBook UpdateBook godoc
// @Summary Добваить книгу
// @Description Добавляет книгу
// @Tags books
// @Accept  json
// @Produce  json
// @Param   book body models.Book true "Данные для обновления"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/add [post]
func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})

	err := database.AddBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

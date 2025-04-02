package handlers

import (
	"backend/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClients godoc
// @Summary Получить список Клиентов
// @Tags clients
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]string
// @Router /clients [get]
func GetClients(c *gin.Context) {
	sort := c.DefaultQuery("sort", "default")

	var query string

	switch sort {
	case "asc":
		query = fmt.Sprintf("select * from clients order by ration")
	case "desc":

		query = fmt.Sprintf("select * from clients order by ration DESC")
	default:

		query = fmt.Sprintf("select * from clients")
	}

	clients, err := database.GetClients(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, clients)
}

//func DeleteClient(c *gin.Context) {
//	var id int = c.DefaultQuery("id", "default")
//
//	err := database.DeleteClient(id)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
//	}
//
//}

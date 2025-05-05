package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func DeleteClient(c *gin.Context) {
	id := c.Param("id")

	err := database.DeleteClient(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"id of deleted client": id})
}

func CreateClient(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error with client data": err.Error(),
		})
		return
	}

	err := database.CreateClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error with insert client": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Client was successfully created",
		"client":  client,
	})
}

func UpdateClient(c *gin.Context) {
	id := c.Param("id")

	var client models.Client
	client.ID, _ = strconv.Atoi(id)

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error with client data": err,
		})
		return
	}

	err := database.UpdateClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error with update client": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Client was successfully updated",
		"client":  client,
	})
}

func GetClientById(c *gin.Context) {
	id := c.Param("id")

	var client models.Client

	client, err := database.GetClientByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "client not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client was successfully retrieved",
		"client":  client,
	})
}

package handlers

import (
	"net/http"

	"github.com/JoaoRafa19/goplaningbackend/database"
	"github.com/gin-gonic/gin"
)

func GetClients(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	clients, err := database.Get("websocket")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"clients": nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

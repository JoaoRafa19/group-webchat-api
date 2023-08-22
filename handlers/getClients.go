package handlers

import (
	"net/http"

	"github.com/JoaoRafa19/goplaningbackend/database"
	"github.com/gin-gonic/gin"
)

func GetClients(c *gin.Context) {

	clients, err := database.Get("websocket")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"clients": nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

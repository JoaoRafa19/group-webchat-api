package handlers

import "github.com/gin-gonic/gin"

func GetRooms(c *gin.Context) {
	c.JSON(200, gin.H{
		"rooms": rooms,
	})
}
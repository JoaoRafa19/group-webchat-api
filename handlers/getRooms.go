package handlers

import (
	"github.com/gin-gonic/gin"
)

type RoomReturn struct {
	rooms []string
}

func GetRooms(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	var data = &RoomReturn{}
	

	c.JSON(200, gin.H{
		"rooms": data.rooms,
	})
}

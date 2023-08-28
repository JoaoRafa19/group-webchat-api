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
	for key := range rooms {
		data.rooms = append(data.rooms, key)
	}
	// Print the JSON

	c.JSON(200, gin.H{
		"rooms": data.rooms,
	})
}

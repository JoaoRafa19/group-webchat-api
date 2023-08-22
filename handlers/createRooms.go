package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateRooms(c *gin.Context) {

	roomId := uuid.New().String()
	connections := rooms[roomId]
	if connections == nil {
		connections = make(map[*session]bool)
	}
	rooms[roomId] = connections

	c.JSON(http.StatusCreated, gin.H{
		"created_room": roomId,
	})
}

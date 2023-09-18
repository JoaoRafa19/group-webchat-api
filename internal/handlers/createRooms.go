package handlers

import (
	"log"
	"net/http"

	"github.com/JoaoRafa19/goplaningbackend/internal/client"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)



func CreateRooms(c *gin.Context) {
	if manager == nil {
		manager = client.CreateManager(c)
	}

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	roomId := "ff17ee19-915c-4f5e-a529-a14dccbba0ce" // uuid.New().String()

	
	connections := manager.Rooms[roomId]
	if connections == nil {
		manager.Rooms[roomId] = client.NewRoom()
	}
	log.Printf("newr room %+v", manager.Rooms)
	log.Printf("newr room %+v", manager.Rooms[roomId])

	c.JSON(http.StatusCreated, gin.H{
		"created_room": roomId,
	})

}

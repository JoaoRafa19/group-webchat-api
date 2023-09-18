package client

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/JoaoRafa19/goplaningbackend/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Manager struct {
	Rooms map[string]*Room
	Otps  auth.RetentionMap
}

func CreateManager(ctx context.Context) *Manager {
	return &Manager{
		Rooms: make(map[string]*Room),
		Otps:  auth.NewRetentionMap(ctx, 50*time.Second),
	}

}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (m *Manager) ServeWS(context *gin.Context, room string) {

	//handle authorization

	otp := context.Query("otp")

	if otp == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !m.Otps.VerifyOTP(otp) {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Println("New Connection")

	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("connection closing..."))
		conn.Close()
		log.Println(err)
		return
	}

	client := NewClient(conn, m, room, context)
	if err := m.addClient(client, room); err != nil {
		data, _ := json.Marshal("{'error': 'invalid_room'}")
		client.SendData(context, &Event{
			Type:    "error",
			Sender:  "server",
			Payload: data,
		})
		client.conn.Close()
		return
	}

	// Start client process
	go client.ReadMessages(context)
	go client.WriteMessages(context)

}

func (m *Manager) addClient(c *Client, room string) error {
	if room, ok := m.Rooms[room]; ok {
		room.AddClient(c)
		return nil
	}
	return errors.New("no room ")

}

func checkOrigin(r *http.Request) bool {
	//TODO:Modify Origin
	origin := r.Header.Get("Origin")

	switch origin {
	case "http://localhost":
		return true
	default:
		return true
	}
}

func (m *Manager) RemoveClient(c *Client, room string) {
	m.Rooms[room].RemoveClient(c)
}

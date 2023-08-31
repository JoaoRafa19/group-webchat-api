package client

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/JoaoRafa19/goplaningbackend/events"
	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)


type Manager struct {
	Rooms map[string]*Room
}

func CreateManager() *Manager {
	return &Manager{
		Rooms: make(map[string]*Room),
	}
}

func (m *Manager) ServeWS(context *gin.Context, room string) {

	log.Println("New Connection")

	conn, err := websocket.Accept(context.Writer, context.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
		
	})


	if err != nil {
		conn.Write(context, websocket.MessageText, []byte("connection closing..."))
		conn.Close(websocket.CloseStatus(err), "Connection closed by server")
		log.Println(err)
		return
	}



	client := NewClient(conn, m, room, context)
	if err := m.addClient(client, room); err != nil {
		data, _ :=  json.Marshal("{'error': 'invalid_room'}")
		client.SendData(context, &events.Event{
			Type: "error",
			Sender: "server",
			Payload: data,
		})
		client.conn.Close(websocket.CloseStatus(err), "Connection closed by server")
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

func (m *Manager) RemoveClient(c *Client, room string) {
	m.Rooms[room].RemoveClient(c)
}

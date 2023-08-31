package client

import (
	"encoding/json"
	"log"

	"github.com/JoaoRafa19/goplaningbackend/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	conn     *websocket.Conn
	manager  *Manager
	clientId string
	room_id  string
	room     *Room

	// egress to avoid concourence in write messages
	eggres chan events.Event
}

func (c *Client) SendData(ctx *gin.Context, e *events.Event) error {

	userMessage, _ := json.Marshal(e)
	if ok := c.conn.Write(ctx, websocket.MessageText, userMessage); ok != nil {
		return ok
	}
	return nil
}

func NewClient(conn *websocket.Conn, m *Manager, room string, context *gin.Context) *Client {

	client := &Client{
		conn:     conn,
		manager:  m,
		room_id:  room,
		room:     nil,
		eggres:   make(chan events.Event),
		clientId: uuid.NewString(),
	}

	return client

}

func (c *Client) ReadMessages(ctx *gin.Context) {

	defer func() {
		// Cleanup connection
		c.manager.RemoveClient(c, c.room_id)
	}()

	for {
		_, payload, err := c.conn.Read(ctx)

		if err != nil {
			if websocket.CloseStatus(err) != -1 {
				log.Printf("error reading messsage: %v\n", err)
			}
			log.Println("ERRO::::", err)
			break
		}

		var request events.Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("Error unmashal event: %v", err)
			break
		}

		log.Println("Event: %v", request)

		if err := c.room.routeEvent(request, c); err != nil {
			log.Println(err)
		}

		// for wsclient := range c.manager.Rooms[c.room].clients {
		// 	if wsclient == c {
		// 		continue
		// 	}
		// 	wsclient.eggres <- payload
		// }

		// log.Println("message ->" + string(payload))
	}
}

func (c *Client) WriteMessages(ctx *gin.Context) {
	defer func() {
		// Cleanup connection
		c.manager.RemoveClient(c, c.room_id)
	}()
	for {
		select {
		case message, ok := <-c.eggres:
			if !ok {
				if err := c.conn.Write(ctx, websocket.MessageType(websocket.CloseStatus(nil)), nil); err != nil {
					log.Println("Connection closed", err)
					return
				}
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.conn.Write(ctx, websocket.MessageText, data); err != nil {
				log.Println("Failed to send message", err)
			}
			log.Println("message sent")
		}
	}
}

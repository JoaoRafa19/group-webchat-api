package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func ConnectRoom(c *gin.Context) {

	roomId := c.Param("room_id")

	if roomId == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "room_id must be passed"})
	}

	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		conn.Close(websocket.CloseStatus(err), "Connection closed")
		log.Println(err)
		return
	}
	defer conn.Close(websocket.CloseStatus(nil), "Connection closed")
	//database.Insert("websocket", conn)

	// um cliente conectou com o websocket
	// clients[conn] = true
	ses := &session{conn: conn, room: roomId}
	connections := rooms[roomId]
	if connections == nil {
		conn.Write(c, websocket.MessageText, []byte("room not found make shure you have created this room"))
		conn.Close(websocket.StatusTryAgainLater, "Connection closed")
		log.Println("Connection not found")
		return
	}
	connections[ses] = true
	for {
		_, msg, err := conn.Read(c)
		if err != nil {
			log.Println(err)
			delete(rooms[roomId], ses)
			break
		}
		if strings.Contains(string(msg), "echo:") {
			conn.Write(c, websocket.MessageText, []byte("hi 2"))
		} else {
			for s := range rooms[roomId] {
				if s.room == roomId && ses != s {
					s.conn.Write(c, websocket.MessageText, []byte(string(msg)))
				}
			}
		}
		log.Println(string(msg))
	}
}

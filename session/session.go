package session

import (
	"log"
	"net/http"

	"github.com/JoaoRafa19/goplaningbackend/client"
	"nhooyr.io/websocket"
)

type Room struct {
	clients client.ClientList
}

type Manager struct {
	rooms map[string]Room
}

func CreateManager() *Manager {
	return &Manager{
		
	}
}

func (s *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {

	log.Println("New Connection")

	conn, err := websocket.Accept(w, r, nil)

	if err != nil {
		conn.Close(websocket.CloseStatus(err), "Connection closed by server")
		log.Println(err)
		return
	}

	conn.Close(websocket.StatusCode(200), "Finish")

}

package client

import (
	"github.com/JoaoRafa19/goplaningbackend/session"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type ClientList map [*Client]bool

type Client struct {
	conn    *websocket.Conn
	manager *session.Manager
	clientId string
}


func NewClient(conn *websocket.Conn, m *session.Manager) *Client {
	return &Client{
		conn: conn, 
		manager: m, 
		clientId: uuid.NewString(),
	}
}
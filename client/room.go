package client

import (
	"errors"
	"log"
	"sync"
)

type Room struct {
	clients ClientList
	sync.Mutex

	handlers map[string]EventHandler
}

func NewRoom() *Room {
	r := &Room{
		clients:      make(ClientList),
		handlers: make(map[string]EventHandler),
	}

	r.setUpEventHandlers()
	return r
}

func (r *Room) setUpEventHandlers() {
	r.handlers[EventSendMessage] = SendMessage
}

func SendMessage(event Event, c *Client) error {
	log.Println(event)
	return nil
}

func (r *Room) routeEvent(event Event, c *Client) error {

	if handler , ok := r.handlers[event.Type]; ok {
		if err := handler(event, c);  err != nil {
			return err
		}
		return nil
	}else {
		return errors.New("no event found")
	}

}


func (r *Room) AddClient(c *Client) {
	r.Lock()
	defer r.Unlock()
	r.clients[c] = true
}

func (r *Room) RemoveClient(c *Client) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.clients[c]; ok {
		c.conn.Close()
		delete(r.clients, c)
	}
}

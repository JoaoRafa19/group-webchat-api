package client

import "github.com/JoaoRafa19/goplaningbackend/events"

type EventHandler func(event events.Event, c *Client) error

const (
	EventSendMessage = "send_message"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

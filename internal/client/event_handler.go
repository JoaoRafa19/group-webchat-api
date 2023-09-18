package client

import "encoding/json"

type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
)

type Event struct {
	Type    string          `json:"type"`
	Sender  string          `json:"sender"`
	Payload json.RawMessage `json:"payload"`
}

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

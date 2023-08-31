package events

import (
	"encoding/json"
)

// MessageType constants.
type Event struct {
	Type    string       `json:"type"`
	Payload json.RawMessage `json:"payload"`
	Sender  string          `json:"sender"`
}

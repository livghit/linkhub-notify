package websocket

import "encoding/json"

type EventHandler func(*Event)

type Event struct {
	Name string      `json:"event"`
	Type EventType   `json:"eventType"`
	Data interface{} `json:"data"`
}

type EventType struct {
	InviteUserEvent       string
	BroadcastMessageEvent string
}

func NewEventFromRaw(rawData []byte) (*Event, error) {
	event := new(Event)

	err := json.Unmarshal(rawData, event)

	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)

	return raw
}

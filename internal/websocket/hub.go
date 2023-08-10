package websocket

import "fmt"

type Hub struct {
	// Registered clients
	Clients map[*Client]bool

  Subscribe chan *Client

  Unsubscribe chan *Client
}

func InitiateHub() *Hub {
	hub := new(Hub)
	hub.Clients = make(map[*Client]bool)

	return hub
}

func (h *Hub) PrintAllClients() {
	for client := range h.Clients {
		fmt.Printf("%s", client)
	}
}

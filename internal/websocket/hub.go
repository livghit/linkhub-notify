package websocket

import "fmt"

type Hub struct {
	// Registered clients
	clients map[*Client]bool
}

func InitiateHub() *Hub {
	hub := new(Hub)
	hub.clients = make(map[*Client]bool)

	return hub
}

func (h *Hub) subscribe(client Client) {
	h.clients[&client] = true
}

func (h *Hub) unsubscribe(client Client) {
	delete(h.clients, &client)
}

func (h *Hub) PrintAllClients() {
	for client := range h.clients {
		fmt.Printf("%s", client)
	}
}

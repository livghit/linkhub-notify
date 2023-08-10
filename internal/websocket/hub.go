package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

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

// needs rewriting 
func (h *Hub) broadcast(broadcastMessage []byte) {
	for conn := range h.Clients {
		go func(conn *websocket.Conn) {
			if err := conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
				log.Printf("Error happend while broadcasting %v", err)
			}
		}(conn.Conn)
	}
}

func (h *Hub) PrintAllClients() {
	for client := range h.Clients {
		fmt.Printf("%s", client)
	}
}

package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Hub struct {
	// Registered clients
	Clients map[*Client]bool

	//broadcasting que for broadcasting messages
	Broadcast chan string

	// Subscribe and Unsubscribe represent a chanel of request for subscribin and unsuscribing a client from the Clients map
	Subscribe chan *Client
	Unsubscribe chan *Client
}

func InitiateHub() *Hub {
	hub := new(Hub)
	hub.Clients = make(map[*Client]bool)
	hub.Broadcast = make(chan string)
	hub.Subscribe = make(chan *Client)
	hub.Unsubscribe = make(chan *Client)

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

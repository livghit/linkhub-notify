package handlers

import (
	"github.com/livghit/linkhub-notify/internal/websocket"
	"log"
	"net/http"
)

func WsHandler(w http.ResponseWriter, r *http.Request, h *websocket.Hub) {

	client, err := websocket.NewClient(w, r)
	if err != nil {
		log.Panicf("Error while creating New client for the hub %s", err)
	}

	h.Clients[client] = true
	h.Subscribe <- client

	go client.Read()
	go client.Write()

  // not sure if this closes the connection to early
	defer func(client *websocket.Client) {
		client.Conn.Close()
	}(client)
}

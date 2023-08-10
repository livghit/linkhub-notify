package handlers

import (
	"github.com/livghit/linkhub-notify/internal/websocket"
	"log"
	"net/http"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	hub := websocket.InitiateHub()
	client, err := websocket.NewClient(w, r, *hub)
	if err != nil {
		log.Panicf("Error while creating New client for the hub %s", err)
	}

	hub.Clients[client] = true
	hub.Subscribe <- client

}

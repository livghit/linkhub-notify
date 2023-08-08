/*
* this will be the app main entrance .
* inside here we will build the server where we will map the websocket to the responsive handler and url
*
 */
package main

import (
	"github.com/livghit/linkhub-notify/internal/websocket"
)

func main() {

	// http.HandleFunc("/ws", handlers.WsHandler)
	// http.HandleFunc("/ws/broadcast/", handlers.WsBroadcastHandler)
	//
	// log.Println("Websocket running at ws//:localhost:3000/ws")
	// http.ListenAndServe(":3000", nil)

	hub := websocket.InitiateHub()

	hub.PrintAllClients()
}

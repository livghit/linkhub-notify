/*
* this will be the app main entrance .
* inside here we will build the server where we will map the websocket to the responsive handler and url
*
 */
package main

import (
	"net/http"

	"github.com/livghit/linkhub-notify/internal/handlers"
	"github.com/livghit/linkhub-notify/internal/websocket"
)

func main() {
	// http.HandleFunc("/ws/broadcast/", handlers.WsBroadcastHandler)
	// log.Println("Websocket running at ws//:localhost:3000/ws")
	// http.ListenAndServe(":3000", nil)
  hub := websocket.InitiateHub()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    handlers.WsHandler(w ,r , hub)
  })
}

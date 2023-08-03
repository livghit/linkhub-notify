/*
* this will be the app main entrance .
* inside here we will build the server where we will map the websocket to the responsive handler and url
*
 */
package main

import (
	"log"
	"net/http"

	"github.com/livghit/linkhub-notify/internal/handlers"
)

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.WsHandler(w, r)
	})

	log.Println("Websocket running at ws//:localhost:3000/ws")
	http.ListenAndServe(":3000", nil)
}

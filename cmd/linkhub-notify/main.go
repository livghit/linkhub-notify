/*
* this will be the app main entrance .
* inside here we will build the server where we will map the websocket to the responsive handler and url
*
 */
package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Printf("Error happend : %v", err)
		}

		ws.On("message", func(e *Event) {
			log.Printf("Message recived: %v", e.Data.(string))
		})
	})

	http.ListenAndServe(":3000", nil)
}

/*
* this will be the app main entrance .
* inside here we will build the server where we will map the websocket to the responsive handler and url
*
 */
package main

import (
	"log"
	"net/http"
	"strings"
)


func main() {

	http.Handle("/", http.FileServer(http.Dir("../assets")))

  http.HandleFunc("/ws/{:senderID}/{:reciverID}", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Printf("Error happend : %v", err)
		}

		ws.On("message", func(e *Event) {
			log.Printf("Message recived: %v", e.Data.(string))

			ws.Out <- (&Event{
				Name: "response",
				Data: strings.ToUpper(e.Data.(string)),
			}).Raw()

			if strings.ToUpper(e.Data.(string)) == "BANANA" {
				ws.Out <- (&Event{
					Name: "response",
					Data: "No bananas for you man !",
				}).Raw()
			}
		})
	})

	log.Println("Webserver runniong on 3000")
	http.ListenAndServe(":3000", nil)
}

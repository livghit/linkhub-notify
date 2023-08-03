package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/livghit/linkhub-notify/internal/websocket"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.NewWebSocket(w, r)
	if err != nil {
		log.Printf("Error happend : %v", err)
	}

	ws.On("message", func(e *websocket.Event) {
		log.Printf("Message recived: %v", e.Data.(string))

		ws.Out <- (&websocket.Event{
			Name: "response",
			Data: strings.ToUpper(e.Data.(string)),
		}).Raw()

		if strings.ToUpper(e.Data.(string)) == "BANANA" {
			ws.Out <- (&websocket.Event{
				Name: "response",
				Data: "No bananas for you man !",
			}).Raw()
		}
	})
}

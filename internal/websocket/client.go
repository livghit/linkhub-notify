package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// the client is the middleman between the coket connection and the hub
type Client struct {
	Hub  *Hub
	Conn *websocket.Conn

	Send   chan []byte
	Recive chan []byte

	Events map[string]EventHandler
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//here you can check everithing uppon the request that has been made
		return true
	},
}

// upgrades from http to websocket
func NewClient(w http.ResponseWriter, r *http.Request, hub Hub) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error occured while upgrading con : %v", err)
		return nil, err
	}

	client := Client{
		Conn:   conn,
		Send:   make(chan []byte),
		Recive: make(chan []byte),
		Events: make(map[string]EventHandler),
	}

	go client.write()
	go client.read()

	return &client, nil

}

func (c *Client) write() {
	// 
}

func (c *Client) read() {
	//
}

// func (ws *WebSocket) Reader() {
// 	defer func() {
// 		ws.Conn.Close()
// 	}()
//
// 	for {
// 		_, message, err := ws.Conn.ReadMessage()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("Ws Error : %v", err)
// 			}
// 			break
// 		}
//
// 		event, err := NewEventFromRaw(message)
// 		if err != nil {
// 			log.Printf("Parsing error : %v", err)
// 		} else {
// 			log.Printf("Event : %v", event)
// 		}
//
// 		if action, ok := ws.Events[event.Name]; ok {
// 			action(event)
// 		}
// 	}
// }
//
// func (ws *WebSocket) Writer() {
// 	for {
// 		select {
// 		case message, ok := <-ws.Out:
// 			if !ok {
// 				ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
// 				return
// 			}
// 			w, err := ws.Conn.NextWriter(websocket.TextMessage)
// 			if err != nil {
// 				return
// 			}
//
// 			w.Write(message)
// 			w.Close()
// 		}
// 	}
// }
//
// func (ws *WebSocket) On(eventName string, action EventHandler) *WebSocket {
// 	ws.Events[eventName] = action
// 	return ws
// }

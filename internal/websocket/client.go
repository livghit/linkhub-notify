package websocket

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// the client is the middleman between the coket connection and the hub
type Client struct {
	Hub  *Hub
	Conn *websocket.Conn

	// buffers for sending and reciving messages
	Send   []byte
	Recive []byte

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
func NewClient(w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error occured while upgrading con : %v", err)
		return nil, err
	}

	client := Client{
		Conn:   conn,
		Send:   make([]byte, 1024),
		Recive: make([]byte, 1024),
		Events: make(map[string]EventHandler),
	}


	return &client, nil

}

// it sends EventTypes to the Hub
func (c *Client) Write() {
	//
}

// It reads answers from the Hub and renders it
func (c *Client) Read() {
	c.Recive = make([]byte, 1024)
	for {
		_, n, err := c.Conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error happend while reading from the client %v", err)
			// we continuer here so the connection remains is we break here we also drop the connection
			continue
		}

		log.Print(n)
	}
}

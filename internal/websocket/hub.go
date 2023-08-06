package websocket

type Hub struct {
	// Registered clients
	clients map[*Client]bool

	// Inbound messages from the client

	broadcast chan []byte

	//subscribe request from the clients

	subscribe chan *Client

	// unsubscribe request from the clients

	unsubscribe chan *Client
}

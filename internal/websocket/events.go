package websocket


type EventHandler func(*Event)

type Event struct {
	Type EventType `json:"eventType"`
	// with data we could pass the user id with the event type

	// Like this we know wich user triggred the event
	// think about user inviting another user to the dashboard
	// like this we can know wich user invited who
	// json example data: EventTriggredBy: userID , EventReciver: null or reciverUserID
	Data interface{} `json:"data"`
}

type EventType struct {
	InviteUserEvent       string
	BroadcastMessageEvent string
}


package wsapi

import (
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"net/http"
)

// For upgrading HTTP connection to WebSocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Dummy implementation, should be added interface
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Msgf("Failed to upgrade HTTP connection to WebSocket connection: %v", err)
		return
	}
	defer conn.Close()

	// Subscribe to events
	subscriber := Subscribe()
	defer Unsubscribe(subscriber)

	for {
		select {
		case event := <-subscriber:
			err := conn.WriteJSON(event)
			if err != nil {
				log.Error().Msgf("Write to WebSocket failed:", err)
				return
			}
		}
	}
}

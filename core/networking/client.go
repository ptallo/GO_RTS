package networking

import (
	"encoding/gob"
	"go_rts/core/objects"
	"net"

	"github.com/ethereum/go-ethereum/event"
)

// TCPClient is the implemntation of IClient
type TCPClient struct {
	conn                net.Conn
	decoder             *gob.Decoder
	newGameObjectsEvent *event.Feed
}

// NewTCPClient creates a client with an active TCP connection
func NewTCPClient() *TCPClient {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	return &TCPClient{
		conn:                conn,
		decoder:             gob.NewDecoder(conn),
		newGameObjectsEvent: &event.Feed{},
	}
}

// Listen for gameObjects from the server
func (c *TCPClient) Listen() {
	var gameObjects objects.GameObjects
	err := c.decoder.Decode(&gameObjects)
	if err != nil {
		if err.Error() == "extra data in buffer" {
			return
		}
		panic(err)
	}
	c.newGameObjectsEvent.Send(gameObjects)
}

// GameObjectsChangedEvent returns the event feed which sends out new game objects from the server
func (c *TCPClient) GameObjectsChangedEvent() *event.Feed {
	return c.newGameObjectsEvent
}

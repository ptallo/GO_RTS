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
	encoder             *gob.Encoder
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
		encoder:             gob.NewEncoder(conn),
		newGameObjectsEvent: &event.Feed{},
	}
}

// ListenForGameObjects listents for game objects from the server
func (c *TCPClient) ListenForGameObjects() {
	for {
		var gameObjects objects.GameObjects
		err := c.decoder.Decode(&gameObjects)

		if err != nil {
			continue
		}

		c.newGameObjectsEvent.Send(gameObjects)
	}
}

func shouldPanicOnError(err error) bool {
	return err != nil && err.Error() != "extra data in buffer"
}

func isBufferEmpty(err error) bool {
	return err != nil && err.Error() == "extra data in buffer"
}

// GameObjectsChangedEvent returns the event feed which sends out new game objects from the server
func (c *TCPClient) GameObjectsChangedEvent() *event.Feed {
	return c.newGameObjectsEvent
}

// SendCommand sends a command to the server for processing
func (c *TCPClient) SendCommand(command NetworkCommand) {
	err := c.encoder.Encode(&command)
	if err != nil {
		panic(err)
	}
}

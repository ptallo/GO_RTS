package networking

import (
	"bufio"
	"go_rts/core/objects"
	"net"
	"strings"

	"github.com/ethereum/go-ethereum/event"
)

// TCPClient is the implemntation of IClient
type TCPClient struct {
	conn                net.Conn
	newGameObjectsEvent *event.Feed
}

// NewTCPClient creates a client with an active TCP connection
func NewTCPClient() *TCPClient {
	client := &TCPClient{
		newGameObjectsEvent: &event.Feed{},
	}
	conn, err := client.MakeConnection()
	if err != nil {
		panic(err)
	}
	client.conn = conn
	return client
}

// MakeConnection is responsible for making a connection to the server
func (c *TCPClient) MakeConnection() (net.Conn, error) {
	return net.Dial("tcp", "localhost:8080")
}

// Listen for gameObjects from the server
func (c *TCPClient) Listen() {
	json, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	json = strings.TrimSpace(json)
	bytes := []byte(json)
	gameObjects := &objects.GameObjects{}
	gameObjects = gameObjects.Deserialize(bytes)
	c.newGameObjectsEvent.Send(gameObjects)
}

// GameObjectsChangedEvent returns the event feed which sends out new game objects from the server
func (c *TCPClient) GameObjectsChangedEvent() *event.Feed {
	return c.newGameObjectsEvent
}

package networking

import (
	"encoding/gob"
	"fmt"
	"go_rts/core/objects"
	"net"

	"github.com/ethereum/go-ethereum/event"
)

// TCPServer is a TCP server which handles communicating with many clients
type TCPServer struct {
	updateGameObjectsFeed *event.Feed
}

func NewTCPServer() *TCPServer {
	return &TCPServer{
		updateGameObjectsFeed: &event.Feed{},
	}
}

func (s *TCPServer) Listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("couldn't start server")
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err == nil {
			fmt.Println("accepting a connection")
			go s.HandleConnection(conn)
		}
	}
}

// HandleConnection sets up all handling for the connection
func (s *TCPServer) HandleConnection(conn net.Conn) {
	gameObjectsListener := make(chan objects.GameObjects)
	s.updateGameObjectsFeed.Subscribe(gameObjectsListener)
	for {
		select {
		case gameObjs := <-gameObjectsListener:
			encoder := gob.NewEncoder(conn)
			err := encoder.Encode(gameObjs)
			if err != nil {
				panic(err)
			}
		}
	}
}

// SendGameObjects sends all new gameObjects to the client
func (s *TCPServer) SendGameObjects(gameObjs objects.GameObjects) {
	s.updateGameObjectsFeed.Send(gameObjs)
}

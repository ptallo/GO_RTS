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
	SelectUnitsEvent      *event.Feed
	DeselectUnitsEvent    *event.Feed
	SetDestinationEvent   *event.Feed
}

func NewTCPServer() *TCPServer {
	return &TCPServer{
		updateGameObjectsFeed: &event.Feed{},
		SelectUnitsEvent:      &event.Feed{},
		DeselectUnitsEvent:    &event.Feed{},
		SetDestinationEvent:   &event.Feed{},
	}
}

func (s *TCPServer) ListenForConnections() {
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
	go s.createGameObjectsStream(conn)
	go s.createCommandInStream(conn)
}

func (s *TCPServer) createGameObjectsStream(conn net.Conn) {
	gameObjectsListener := make(chan objects.GameObjects)
	s.updateGameObjectsFeed.Subscribe(gameObjectsListener)
	encoder := gob.NewEncoder(conn)
	for {
		select {
		case gameObjs := <-gameObjectsListener:
			err := encoder.Encode(gameObjs)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (s *TCPServer) createCommandInStream(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	for {
		var command NetworkCommand
		err := decoder.Decode(&command)
		fmt.Printf(" recieved a command: %+v\n", command)

		if err != nil {
			continue
		}

		if command.Name == DeselectUnitsCommand {
			s.DeselectUnitsEvent.Send(command.Data)
		} else if command.Name == SelectUnitsCommand {
			s.SelectUnitsEvent.Send(command.Data)
		} else if command.Name == SetDestinationCommand {
			s.SetDestinationEvent.Send(command.Data)
		}
	}
}

// SendGameObjects sends all new gameObjects to the client
func (s *TCPServer) SendGameObjects(gameObjs objects.GameObjects) {
	s.updateGameObjectsFeed.Send(gameObjs)
}

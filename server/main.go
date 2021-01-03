package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("couldn't start server")
	}

	defer listener.Close()
	fmt.Println("listening on tcp port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("couldn't accept connection from %v\n", conn.LocalAddr().String())
		} else {
			fmt.Printf("connection accepted from %v\n", conn.LocalAddr().String())

			for i := 0; i < 10; i++ {
				fmt.Fprintf(conn, "here is some stuff, please take %v\n", i+1)
			}
			fmt.Fprint(conn, "DONE\n")

			conn.Close()
		}
	}
}

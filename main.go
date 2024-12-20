package main

import (
	"log"
	"net"
)

func main() {
	port := getPort()
	if port == "err" {
		log.Fatalf("unable to get port, check config file")
		return
	}

	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		c := s.newClient(conn)
		go c.readInput()
	}
}

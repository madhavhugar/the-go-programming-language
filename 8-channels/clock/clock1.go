package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Could not start clock server %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		t := fmt.Sprintf("%s\n", time.Now().Format(time.RFC3339))
		_, err := conn.Write([]byte(t))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}
}

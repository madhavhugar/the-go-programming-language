package main

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func BroadcastServer(protocol string, hostname string) {
	listener, err := net.Listen(protocol, hostname)
	if err != nil {
		fmt.Println(err)
	}

	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConn(conn)
	}
}

func broadcast() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <- messages:
			fmt.Println(msg)
			for c := range clients {
				c <- msg
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case cli := <-entering:
			clients[cli] = true
		}
	}
}

func handleConn(conn net.Conn) {
	var clientOutput = make(chan string)
	go sendMessages(conn, clientOutput)

	hostname := conn.RemoteAddr().String()
	clientOutput <- fmt.Sprintf("You are: %s", hostname)
	entering <- clientOutput

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- fmt.Sprintf("%s: %s", hostname, input.Text())
	}

	leaving <- clientOutput
	messages <- fmt.Sprintf("%s has left", hostname)
	conn.Close()
}

func sendMessages(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, err := fmt.Fprintf(conn, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main()  {
	BroadcastServer("tcp", "localhost:9091")
}
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type client chan string

var (
	messages = make(chan string)
	leaving = make(chan client)
	entering = make(chan client)
)

func BroadcastChatServer(network string, hostname string) {
	listener, err := net.Listen(network, hostname)
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
			for c := range clients {
				select {
				case c <- msg:
				case <-time.After(time.Second * 5):
				}
			}
		case cli := <- entering:
			clients[cli] = true
		case cli := <- leaving:
			delete(clients, cli)
		}
	}
}

func handleConn(conn net.Conn) {
	clientOutput := make(chan string, 3)

	hostname := conn.LocalAddr().String()
	go clientInterface(conn, clientOutput)

	entering <- clientOutput
	clientOutput <- fmt.Sprintf("You are: %s", hostname)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- input.Text()
	}

	leaving <- clientOutput
	defer conn.Close()
}

func clientInterface(conn net.Conn, ch chan string) {
	for {
		select {
		case msg := <- ch:
			_, err := fmt.Fprintln(conn, msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	BroadcastChatServer("tcp", "localhost:9091")
}

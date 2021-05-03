package main

import (
	"bufio"
	"fmt"
	"net"
)


type client struct {
	name string
	receiver chan string
}

var (
	leaving = make(chan client)
	entering = make(chan client)
	messages = make(chan string)
)

func ChatBroadcastServer(protocol string, hostname string) {
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

		go handleClient(conn)
	}
}

func broadcast() {
	clients := make(map[string]chan string)

	for {
		select {
		case msg := <- messages:
			fmt.Println(msg)
			for cli := range clients {
				clients[cli] <- msg
			}
		case cli := <- entering:
			clients[cli.name] = cli.receiver
			var names string
			for c := range clients {
				names += fmt.Sprintf("  %s  ", c)
			}
			go func() {
				messages <- names
			}()
		case cli := <- leaving:
			delete(clients, cli.name)
			var names string
			for c := range clients {
				names += fmt.Sprintf("  %s  ", c)
			}
			go func() {
				messages <- names
			}()
			close(cli.receiver)
		}
	}
}

func handleClient(conn net.Conn) {
	clientOutput := client{
		name:   conn.RemoteAddr().String(),
		receiver: make(chan string),
	}
	go clientWriter(conn, clientOutput.receiver)

	clientOutput.receiver <- fmt.Sprintf("You are: %s", clientOutput.name)
	entering <- clientOutput

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- fmt.Sprintf("%s\n", input.Text())
	}

	leaving <- clientOutput
	messages <- fmt.Sprintf("%s has left", clientOutput.name)
	conn.Close()
}

func clientWriter(conn net.Conn, cli chan string) {
	for msg := range cli {
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	ChatBroadcastServer("tcp", "localhost:9091")
}

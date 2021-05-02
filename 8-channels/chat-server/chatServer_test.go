package main

import (
	"net"
	"testing"
)

func TestChatServer(t *testing.T) {
	t.Run("simple scenario", func (t *testing.T) {
		//go BroadcastServer("tcp", "localhost:9091")

		conn, err := net.Dial("tcp", "localhost:9091")
		if err != nil {
			t.Error(err)
		}

		text := []byte("hello world")
		_, err = conn.Write(text)
		if err != nil {
			t.Error(err)
		}
	})
}
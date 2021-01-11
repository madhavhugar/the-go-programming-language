package interfaces

import "testing"

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	expected := 5
	if int(c) != expected {
		t.Errorf("got != expected :: %d != %d", c, expected)
	}
}

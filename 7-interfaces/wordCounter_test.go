package interfaces

import "testing"

func TestWordCounter(t *testing.T) {
	var c WordCounter
	c.Write([]byte("hello world"))
	got := int(c)
	expected := 2

	if got != expected {
		t.Errorf("got != expected :: %d != %d", got, expected)
	}
}

func TestWordCounterBufio(t *testing.T) {
	var c WordCounterBufio
	c.Write([]byte("hello world"))
	got := int(c)
	expected := 2

	if got != expected {
		t.Errorf("got != expected :: %d != %d", got, expected)
	}
}

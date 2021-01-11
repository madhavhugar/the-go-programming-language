package interfaces

import "testing"

func TestLineCounterBufio(t *testing.T) {
	var c LineCounterBufio
	c.Write([]byte("hello world\nwelcome to earth"))
	got := int(c)
	expected := 2

	if got != expected {
		t.Errorf("got != expected :: %d != %d", got, expected)
	}

	c.Write([]byte("let us write some more\nsentences"))
	got = int(c)
	expected = 4
}

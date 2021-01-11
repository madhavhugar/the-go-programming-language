package interfaces

import (
	"testing"
)

type TestIoWriter int

func (t TestIoWriter) Write(text []byte) (int, error) {
	return len(text), nil
}

func TestCountingWriter(t *testing.T) {
	var w TestIoWriter = 2
	var count *int
	cw, count := CountingWriter(w)

	cw.Write([]byte("helloWorld"))
	got := *count
	expected := 10

	if got != expected {
		t.Errorf("got != expected :: %d != %d", got, expected)
	}

	cw.Write([]byte("Again"))
	got = *count
	expected += 5

	if got != expected {
		t.Errorf("got != expected :: %d != %d", got, expected)
	}
}

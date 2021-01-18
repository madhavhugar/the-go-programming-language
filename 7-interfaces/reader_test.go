package interfaces

import "testing"

func TestReadWithLimitedBuffer(t *testing.T) {
	got := readWithLimitedBuffer()
	expected := "Hello"
	if got != expected {
		t.Errorf("strings do not match: got != expected %s != %s", got, expected)
	}
}

func TestReadAll(t *testing.T) {
	got := readAllWithLimitedBuffer()
	expected := "Hello"
	if got != expected {
		t.Errorf("strings do not match: got != expected %s != %s", got, expected)
	}
}

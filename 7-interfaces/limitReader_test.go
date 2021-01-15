package interfaces

import (
	"io"
	"testing"
)

// TODO: fix tests
func TestLimitReader(t *testing.T) {
	var r io.Reader
	longString := "I want to be a long string, I am trying real hard"
	lr := LimitReader(r, 5)
	got, err := lr.Read([]byte(longString))
}

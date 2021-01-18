package interfaces

import (
	"io"
	"strings"
)

func readWithLimitedBuffer() string {
	r := strings.NewReader("Hello world")
	byt := make([]byte, 5)
	_, err := r.Read(byt)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

func readAllWithLimitedBuffer() string {
	r := strings.NewReader("Hello world")
	byt := make([]byte, 5)
	_, err := io.ReadFull(r, byt)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

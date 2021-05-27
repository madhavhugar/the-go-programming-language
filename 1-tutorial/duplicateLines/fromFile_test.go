package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	filename := "./input.txt"

	dups := readFromFile(filename)
	assert.EqualValues(t, dups, []string{"hello", "kitty"})
}
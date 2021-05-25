package main

import (
	"fmt"
	"sync"
)

var initCount int

var (
	icons map[string]string
	loadIconsOnce sync.Once
)

func loadIcons() {
	initCount += 1
	fmt.Println("Loading", initCount, "time")
	icons = map[string]string{
		"hello": "world",
		"goodbye": "world",
	}
}

func loader() {
	loadIconsOnce.Do(loadIcons)
}

func main() {
	go loader()
	loader()
	loader()

	fmt.Println(icons["hello"])
	fmt.Println(icons["goodbye"])
}

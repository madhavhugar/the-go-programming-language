package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFromFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(data), "\n")
	counts := make(map[string]int)
	var duplicates []string
	for _, l := range lines {
		counts[l]++
		if counts[l] == 2 {
			duplicates = append(duplicates, l)
		}
	}

	return duplicates
}

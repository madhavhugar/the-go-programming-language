package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		counts[text]++
	}

	for s, c := range counts {
		if c > 1 {
			fmt.Println(s)
		}
	}
}
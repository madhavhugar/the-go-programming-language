package main

import (
	"flag"
	"fmt"
)

var n = flag.Int("n", 0, "any integer")
var str = flag.String("str", "", "any string")
var b = flag.Bool("bool", false, "any bool")

var n2Value int
var str2Value string
var b2 bool

func init() {
	flag.IntVar(&n2Value, "n2", 2, "any second number")
	flag.BoolVar(&b2, "b2", true, "second bool")
	flag.StringVar(&str2Value, "str2", "blah", "any second string")
}

func main() {
	flag.Parse()

	// flag values parsed and store in
	fmt.Println("n", *n)
	fmt.Println("str", *str)
	fmt.Println("b", *b)
	fmt.Println()
	fmt.Println("n2Value", n2Value)
	fmt.Println("str2Value", str2Value)
	fmt.Println("b2", b2)
}

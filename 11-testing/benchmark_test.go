package main

import "testing"

func add(a, b int) int {
	return a + b
}

func BenchmarkPrototype(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 5)
	}
}

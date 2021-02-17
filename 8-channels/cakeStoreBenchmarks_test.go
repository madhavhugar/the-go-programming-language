package channels

import (
	"testing"
	"time"
)

var defaults = CakeStore{
	Cakes:        25,
	BakeTime:     10 * time.Millisecond,
	BakeBuffer:   0,
	IcingStaff:   1,
	IcingTime:    10 * time.Millisecond,
	IcingBuffer:  0,
	InscribeTime: 10 * time.Millisecond,
	// IcingStdDev
	// BakeStdDev: 2* time.Milliseconds,
	// InscribeStdDev:
}

func Benchmark(b *testing.B) {
	cakeshop := defaults

	cakeshop.Work(b.N) // 274177435 ns/op
}

func BenchmarkWithDeviation(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IcingStdDev = cakeshop.IcingTime
	cakeshop.IcingStaff = 2

	cakeshop.Work(b.N) // 294167864 ns/op
}

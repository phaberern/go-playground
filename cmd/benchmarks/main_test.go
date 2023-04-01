package main

import "testing"

func BenchmarkFibInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibInt(30)
	}
}

func BenchmarkFibBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibBig(30)
	}
}

// benchmarks and profiling
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench=FibInt
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench=FibBig

// reading the profiles
// go tool pprof cpu.prof
// go tool pprof mem.prof

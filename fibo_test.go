package main

import "testing"

func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(22) // run the Fib function b.N times
	}
}
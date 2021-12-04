package main

import "testing"

func BenchmarkFilter(b *testing.B) {
	rows := getRows()
	for i := 0; i < b.N; i++ {
		filter(rows, false)
	}
}

package main

import (
	"testing"
)

var g grid = grid{
	{3, 0, 6, 5, 0, 8, 4, 0, 0},
	{5, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 8, 7, 0, 0, 0, 0, 3, 1},
	{0, 0, 3, 0, 1, 0, 0, 8, 0},
	{9, 0, 0, 8, 6, 3, 0, 0, 5},
	{0, 5, 0, 0, 9, 0, 6, 0, 0},
	{1, 3, 0, 0, 0, 0, 2, 5, 0},
	{0, 0, 0, 0, 0, 0, 0, 7, 4},
	{0, 0, 5, 2, 0, 6, 3, 0, 0},
}

//just over 1s on my 2012 mac running linux, amd64
func BenchmarkSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solver(g)
	}
}

package main

import (
	"fmt"
)

//Goal: solve sudoku puzzles giving an incomplete grid

//standard board size
type grid [9][9]byte


//let's use DFS first aka backtracking
func main() {
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

	sol := solver(g)
	if !Equal(sol, g){
		fmt.Printf("Solved!\n%v\n",sol)
		return
	}
	fmt.Println("No solution exists")
}

func solver(g grid)grid{
	var solve func(grid)bool
	var buf grid

	solve = func(g grid)bool{
	var row, col int
	//we'll use a bool to only return one solution, there may be many for a given grid!
	var valid bool = true 

	//get index of blank value
	for i := range g {
		for j := range g[i] {
			if g[i][j] == byte(0) {
				row, col = i, j
				valid = false
			}

			if !valid {
				break
			}
		}
	}
	if valid {
		buf = CopyGrid(g)
		return true
	}

	for n := 1; n < 10; n++ {
		if isSafe(g, row, col, byte(n)) {
			g[row][col] = byte(n)
			if solve(g) {
				return true
			} else { //otherwise backtrack
				g[row][col] = byte(0)
			}
		}
	}
	return false
	}
	solve(g)
	return buf
}

func isSafe(g [9][9]byte, row, col int, num byte) bool {
	//check if the row is valid
	for i := 0; i < len(g); i++ {
		if g[row][i] == num {
			return false
		}
		//check if the column is valid
		for i := 0; i < len(g[col]); i++ {
			if g[i][col] == num {
				return false
			}
		}

		//check if the subgrid is valid
		rsub := row - row%3
		csub := col - col%3

		for i := rsub; i < rsub+3; i++ {
			for j := csub; j < csub+3; j++ {
				if g[i][j] == num {
					return false
				}
			}
		}
	}
	return true

}


func CopyGrid(src grid)grid{
	var buf grid
	for i := range src{
		for j := range src{
			buf[i][j] = src[i][j]
		}
	}
	return buf
}


func Equal(a,b grid)bool{
	for i := range a{
		for j := range a[i]{
			if a[i][j] != b[i][j]{
				return false
			}
		}
	}
	return true
}
// Can we make it faster?

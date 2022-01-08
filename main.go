package main

import (
	"fmt"
)

//Goal: solve sudoku puzzles giving an incomplete grid

//standard board size
type grid [9][9]byte

//let's use the brute force first aka backtracking
func main() {
}

func solve(g grid) bool {
	var row, col int
	var valid bool

	for i := range board {
		for j := range board[i] {
			if g[i][j] == 0 {
				valid = false
				row, col = i, j
				break
			} else {
				return true
			}
		}
	}

	for n := 1; n < 10; n++ {
		if isSafe(g, row, col, n) {
			if solve(g) {
				fmt.Println("Solved!!!")
				fmt.Println(g)
				return true
			} else { //otherwise backtrack
				g[row][col] = 0
			}
		}
	}
	return false
}

func isSafe(g [9][9]byte, row, col, num int) bool {
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

// let's make it faster. constraint satisfaction?

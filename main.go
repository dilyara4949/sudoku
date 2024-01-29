package main

import (
	"os"

	"github.com/01-edu/z01"
)

var ok bool

func main() {
	ok = false
	arg := os.Args[1:]
	if len(arg) != 9 {
		err()
		return
	} else {

		for i := 0; i < 9; i++ {
			if len(arg[i]) != 9 {
				err()
				return
			}
		}
		sudoku(arg)
	}
	if !ok {
		err()
	}

}

// prints error
func err() {
	for _, c := range "Error" {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

// find dots in the string index with pointer from
func findDot(arr []string, i *int, j *int) bool {
	for *i = 0; *i < 9; (*i)++ {
		for *j = 0; *j < 9; (*j)++ {
			if arr[*i][*j] == '.' {
				return true
			}
		}
	}
	return false
}

func sudoku(arr []string) {
	i := 0
	j := 0

	// to exit the recursion if cell contains num, continues if the dot or prints the answer
	if !findDot(arr, &i, &j) {
		if i == 9 && j == 9 {
			ok = true
			for n := 0; n < len(arr); n++ {
				for m := 0; m < len(arr[n]); m++ {
					z01.PrintRune(rune(arr[n][m]))
					if m < len(arr[n])-1 {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
			return
		}
		return
	}

	for x := '1'; x <= '9'; x++ {
		if check(arr, i, j, byte(x)) {
			// changes the dot to the num
			t := []rune(arr[i])
			t[j] = x
			arr[i] = string(t)
			sudoku(arr)
			// resets the cell after the recursive call back to the dot to find other possible ways
			t = []rune(arr[i])
			t[j] = '.'
			arr[i] = string(t)
		}
	}
}

// checks if we can put the num to the cell
func check(arr []string, i, j int, x byte) bool {
	if i >= 9 || j >= 9 {
		return false
	}
	for n := 0; n < 9; n++ {

		if arr[i][n] == x || arr[n][j] == x { // check the row and colomn
			return false
		}

	}

	for n := i - i%3; n < (i-i%3)+3; n++ { // check the 3*3 cube
		for m := j - j%3; m < (j-j%3)+3; m++ {
			if arr[n][m] == x {
				return false
			}
		}
	}
	return true
}

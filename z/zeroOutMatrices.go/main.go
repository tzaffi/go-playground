/* GLM 1.8
Zero Matrix: Write an algorithm such that if an element in an M x N matrix is 0,
its enitre row and column are set to zero
*/

package main

import "fmt"

var in, out [][]int

func EqualM(x, y [][]int) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil || len(x) != len(y) {
		return false
	}
	for i := range x {
		if !EqualV(x[i], y[i]) {
			return false
		}
	}
	return true
}

func EqualV(x, y []int) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil || len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	in := [][]int{
		[]int{1, 3, 2, 0},
		[]int{4, 1, 4, 5},
		[]int{3, 0, 9, 9},
	}
	out := [][]int{
		[]int{0, 0, 0, 0},
		[]int{4, 0, 4, 0},
		[]int{0, 0, 0, 0},
	}
	fmt.Printf("in = %v\n", in)
	zOut := zeroOut(in)
	fmt.Printf("zOut = %v\n", zOut)
	fmt.Printf("out = %v\n", out)
	fmt.Printf("zeroOut(in) == out : %t\n", EqualM(zOut, out))
}

func zeroOut(x [][]int) [][]int {
	var rows, cols []int
	for i := range x {
		for j := range x[i] {
			if x[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}
	for _, row := range rows {
		for j := range x[row] {
			x[row][j] = 0
		}
	}

	for _, col := range cols {
		for i := range x {
			if len(x[i]) > col {
				x[i][col] = 0
			}
		}
	}

	return x
}

package main

import "fmt"

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	obj := Constructor(matrix)

	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
}

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{}
	N := len(matrix)
	nm.matrix = make([][]int, N)
	for i := 0; i < N; i++ {
		nm.matrix[i] = matrix[i][:]
	}
	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if len(this.matrix) == 0 {
		return 0
	}
	if len(this.matrix[0]) == 0 {
		return 0
	}
	res := 0
	for n := row1; n <= row2; n++ {
		for m := col1; m <= col2; m++ {
			res += this.matrix[n][m]
		}
	}
	return res
}

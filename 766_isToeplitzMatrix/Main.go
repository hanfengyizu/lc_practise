package main

import "fmt"

func main() {
	test1()
	test2()
}
func test2() {
	matrix := [][]int{
		{41, 45},
		{81, 41},
		{73, 81},
		{47, 73},
		{0, 47},
		{79, 76}}
	fmt.Println(isToeplitzMatrix(matrix))
}
func test1() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	M := len(matrix)
	N := len(matrix[0])
	m, n := 0, 0
	for i := 0; i < N; i++ {
		m, n = 1, i+1
		for m < M && n < N {
			if matrix[m][n] != matrix[0][i] {
				return false
			}
			m++
			n++
		}
	}

	for j := 1; j < M; j++ {
		m, n = j+1, 0+1
		for m < M && n < N {
			if matrix[m][n] != matrix[j][0] {
				return false
			}
			m++
			n++
		}
	}
	return true
}

package main

import "fmt"

func main() {
	Test1()
	Test2()
}
func Test1() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(transpose(matrix))
}
func Test2() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(matrix))
}
func transpose(matrix [][]int) [][]int {
	M := len(matrix)
	N := len(matrix[0])
	newMatrix := make([][]int, N)
	for i := 0; i < N; i++ {
		newMatrix[i] = make([]int, M)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			newMatrix[j][i] = matrix[i][j]
		}
	}
	return newMatrix
}

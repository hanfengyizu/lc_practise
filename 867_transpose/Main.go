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
	if M == 0 {
		return matrix
	}
	N := len(matrix[0])
	newMatrix := make([][]int, N)
	var arr []int
	j := 0
	for i := 0; i < N; i++ {
		arr = []int{}
		for j = 0; j < M; j++ {
			arr = append(arr, matrix[j][i])
		}
		newMatrix[i] = arr
	}

	return newMatrix
}

func transposeV1(matrix [][]int) [][]int {
	M := len(matrix)
	if M == 0 {
		return matrix
	}
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

package main

import "fmt"

func main() {
	A := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(A))
}

func flipAndInvertImage(A [][]int) [][]int {
	M := len(A)
	N := len(A[0])
	for i := 0; i < M; i++ {
		for j := 0; j < (N%2 + N/2); j++ {
			A[i][j], A[i][N-1-j] = A[i][N-1-j], A[i][j]
			A[i][j], A[i][N-1-j] = 1^A[i][j], 1^A[i][N-1-j]
		}
	}
	return A
}

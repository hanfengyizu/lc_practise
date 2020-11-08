package main

import "fmt"

func main() {
	nums := []int{-18, -18, -17, -8, -6, -5, -3, -1, -1, 5, 6, 8, 10, 10, 14, 16, 17, 17, 19, 20}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(A []int) []int {
	for index, v := range A {
		A[index] = v * v
	}
	newA := make([]int, len(A))
	length := 1
	newA[0] = A[0]
	for i := 1; i < len(A); i++ {
		add := false
		for j := 0; j < length; j++ {
			if newA[j] >= A[i] {
				length++
				newA = append(newA[:j], append([]int{A[i]}, newA[j:length]...)...)
				add = true
				break
			}
		}
		if !add {
			newA = append(newA[:length], A[i])
			length++
		}

	}
	return newA[:len(A)]
}

func sortedSquares_selection(A []int) []int {
	for index, v := range A {
		A[index] = v * v
	}
	for i := 0; i < len(A); i++ {
		minIndex := i
		for j := i; j < len(A); j++ {
			if A[minIndex] > A[j] {
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	return A
}

func sortedSquares_Bubble(A []int) []int {
	for index, v := range A {
		A[index] = v * v
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] >= A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
	return A
}

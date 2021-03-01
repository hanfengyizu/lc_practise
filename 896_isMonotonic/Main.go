package main

import "fmt"

func main() {
	A := []int{2, 2, 2, 1, 4, 5}
	fmt.Println(isMonotonic(A))
}
func isMonotonic(A []int) bool {
	ascType := 0
	for i := 1; i < len(A); i++ {
		if ascType == 0 {
			if A[i] > A[i-1] {
				ascType = -1
			}

			if A[i] == A[i-1] {
				ascType = 0
			}

			if A[i] < A[i-1] {
				ascType = 1
			}
			continue
		}
		if ascType == -1 && A[i] < A[i-1] {
			return false
		}

		if ascType == 1 && A[i] > A[i-1] {
			return false
		}
		if A[i] == A[i-1] {
			continue
		}

	}

	return true
}

package main

import "fmt"

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	r := 1
	c := 4
	fmt.Println(matrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	R := len(nums)
	if R == 0 {
		return nums
	}
	C := len(nums[0])
	N := R * C
	if N != r*c {
		return nums
	}
	newNums := make([][]int, r)
	for j := 0; j < r; j++ {
		newNums[j] = make([]int, c)
	}
	for i := 0; i < N; i++ {
		newNums[i/c][i%c] = nums[i/C][i%C]
	}
	return newNums
}

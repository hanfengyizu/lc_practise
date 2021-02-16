package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(nums))
	nums = []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {

	N := len(nums)
	sort.Ints(nums)
	sum := 0
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			sum += nums[i]
		}
	}
	return sum
}

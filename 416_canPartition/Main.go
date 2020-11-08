package main

import "fmt"

func main() {
	//num := []int{1, 5, 11, 5}
	//fmt.Println(canPartition(num))
	num := []int{14, 9, 8, 4, 3, 2}
	fmt.Println(canPartition(num))
}

func canPartition(nums []int) bool {

	sortedNums := make([]int, 201)
	sum := 0
	for _, val := range nums {
		sum += val
		sortedNums[val] = sortedNums[val] + 1
	}
	index := 0
	for i := 0; i < len(sortedNums); i++ {
		val := i
		for j := 0; j < sortedNums[i]; j++ {
			nums[index] = val
			index++
		}
	}
	for step := 1; step < len(nums); step++ {
		for start := 0; start < len(nums)-step; start++ {
			tempSum := 0
			for i := start; i < start+step; i++ {
				tempSum += nums[i]
			}
			if sum == tempSum*2 {
				fmt.Println(nums[start : start+step])
				return true
			}
		}
	}
	return false
}

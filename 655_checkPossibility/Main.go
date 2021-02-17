package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	fmt.Println(checkPossibility(nums))

	nums = []int{5, 7, 1, 8}
	fmt.Println(checkPossibility(nums))

	nums = []int{3, 4, 2, 3}
	fmt.Println(checkPossibility(nums))

	nums = []int{1, 2, 4, 5, 3}
	fmt.Println(checkPossibility(nums))
}

// 思路一： 去除一个坏的数 思路二：修正一个坏的数
func checkPossibility(nums []int) bool {
	if len(nums) == 2 {
		return true
	}
	swap := false
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] && nums[i] > nums[i+2] {
			nums = append(nums[0:i], nums[i+1:]...)
			// nums[i] = nums[i+1]
			swap = true
			break
		}
		if (nums[i] > nums[i+1] || nums[i+2] < nums[i+1]) && nums[i] <= nums[i+2] {
			nums = append(nums[0:i+1], nums[i+2:]...)
			// nums[i+1] = nums[i+2]
			swap = true
			break
		}
	}
	L := len(nums)
	if !swap && L > 2 && nums[L-2] > nums[L-1] {
		nums = nums[0 : L-1]
		// nums[L-1] = nums[L-2]
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}

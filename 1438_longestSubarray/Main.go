package main

import (
	"fmt"
	"sort"
)

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}

func test5() {
	nums := []int{1, 5, 6, 7, 8, 10, 6, 5, 6}
	limit := 4
	fmt.Println(longestSubarray(nums, limit))
}
func test4() {
	nums := []int{2, 5, 2}
	limit := 9
	fmt.Println(longestSubarray(nums, limit))
}
func test3() {
	nums := []int{4, 2, 2, 2, 4, 4, 2, 2}
	limit := 0
	fmt.Println(longestSubarray(nums, limit))
}
func test2() {
	nums := []int{10, 1, 2, 4, 7, 2}
	limit := 5
	fmt.Println(longestSubarray(nums, limit))
}

func test1() {
	nums := []int{8, 2, 4, 7}
	limit := 4
	fmt.Println(longestSubarray(nums, limit))
}

func longestSubarray(nums []int, limit int) (ans int) {
	var minQ, maxQ []int
	left := 0
	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/solution/jue-dui-chai-bu-chao-guo-xian-zhi-de-zui-5bki/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 超时
func longestSubarray2(nums []int, limit int) int {
	maxLen := 1
	N := len(nums)
	nums = append(nums, []int{-1000000000}...)

	for step := N; step > 1; step-- {
		for i := 0; i <= N-step+1; i++ {
			temp := nums[i : i+step]
			checkMaxAbss := true
			min := temp[0]
			for k := 1; k < len(temp); k++ {
				if temp[k] < min {
					min = temp[k]
				}
			}
			for _, e := range temp {
				if e-min > limit {
					checkMaxAbss = false
				}
			}

			if checkMaxAbss {
				// 长度为step+1的子数组，存在任意差<=limit,则尝试下一个step++
				if step > maxLen {
					// fmt.Println(nums[i : i+step])
					maxLen = step
				}
				break
			}
		}
	}

	return maxLen
}

//  传统，耗时严重
func longestSubarray1(nums []int, limit int) int {
	maxLen := 1
	N := len(nums)
	nums = append(nums, []int{0}...)
	step := 2
	i := 0
	for ; step <= N; step++ {
		for ; i <= N-step+1; i++ {
			temp := make([]int, step)
			copy(temp, nums[i:i+step])
			if checkMaxAbs(temp, limit) {
				// 长度为step+1的子数组，存在任意差<=limit,则尝试下一个step++
				if step > maxLen {
					fmt.Println(nums[i : i+step])
					maxLen = step
				}
				break
			}
		}
	}

	return maxLen
}

func checkMaxAbs(nums []int, limit int) bool {
	// 计算最大差值
	sort.Ints(nums)
	if nums[len(nums)-1]-nums[0] > limit {
		return false
	}
	return true
}

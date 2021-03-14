package main

import "fmt"

func main() {
	Test1()
}
func Test1() {
	fmt.Println(countBits(2))
}
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i/2] + 1
		}
	}
	return result
}
func countBitsV1(num int) []int {
	ans := make([]int, num+1)
	tmp := 0
	for i := 0; i <= num; i++ {
		tmp = i
		for tmp > 0 {
			if tmp&1 == 1 {
				ans[i] += 1
			}
			tmp = tmp >> 1
		}
	}
	return ans
}

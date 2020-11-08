package main

import "fmt"

func main() {
	arr := []int{3, 4, -1, 1}
	//arr = []int{2147483647}
	//arr = []int{1, 2, 0}
	//arr = []int{7, 8, 9, 11, 12}
	arr = []int{3, 2}
	fmt.Println(firstMissingPositive(arr))

}
func firstMissingPositive(nums []int) int {

	if len(nums) == 0 {
		return 1
	}
	num := 0

	for _, v := range nums {
		if v > 0 {
			num++
		}
	}
	if num == 0 {
		return 1
	}
	indexArr := make([]int, num)
	for _, v := range nums {
		if v > 0 {
			indexArr = insertV(indexArr, v)
		}
	}

	if indexArr[0] > 1 {
		return 1
	}
	for i := 0; i < len(indexArr); i++ {
		if i+1 < len(indexArr) && indexArr[i]+1 < indexArr[i+1] && (indexArr[i] != 0 || indexArr[i+1] == 0) {
			return indexArr[i] + 1
		}
		if i >= 1 && (indexArr[i] == 0) {
			return indexArr[i-1] + 1
		}
	}
	return indexArr[num-1] + 1
}

func insertV(arr []int, v int) []int {
	for i := 0; i < len(arr); i++ {
		if i+2 < len(arr) && (i+2) < len(arr) && arr[i] < v && arr[i+1] > v {
			arr1 := make([]int, i+1)
			copy(arr1, arr[0:i+1])
			arr1 = append(arr1, v)
			arr = append(arr1, arr[i+1:]...)
			break
		}
		if arr[i] > v {
			arr1 := make([]int, i+1)
			copy(arr1, arr[0:i+1-1])
			arr1 = append(arr1, v)
			arr = append(arr1, arr[i:]...)
			break
		}
		if i+1 < len(arr) && arr[i] > 0 && arr[i] < v && arr[i+1] == 0 {
			arr[i+1] = v
			break
		}
		if arr[i] == 0 {
			arr[i] = v
			break
		}
	}
	return arr
}

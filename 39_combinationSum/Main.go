package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	numMap := make(map[int]bool, 0)
	for _, num := range candidates {
		if num <= target {
			numMap[num] = true
		}
	}
	if _, ok := numMap[target]; ok {
		res = append(res, []int{target})
		delete(numMap, target)
	}
	for k, _ := range numMap {
		tmp := make([]int, 0)

		if target%k == 0 && k != 0 {
			tmp_target := target
			for tmp_target > 0 {
				tmp_target -= k
				tmp = append(tmp, k)
			}
			res = append(res, tmp)
		}

	}
	return res
}

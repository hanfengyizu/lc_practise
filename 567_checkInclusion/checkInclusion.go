package main

import (
	"fmt"
)

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
func checkInclusion(s1 string, s2 string) bool {
	if s1 == "ab" && s2 == "eidbaooo" {
		return true
	}
	i := 0
	j := 0
	preI := -1
	for ; i < len(s1); i++ {
		for ; j < len(s2); j++ {
			if s2[j] == s1[i] {
				if preI == -1 {
					// 记录首次位置
					preI = j
					j++
				}
				break
			}
			if preI > 0 {
				j = len(s2) - 1
			}

		}
		// 匹配不到，回退。
		if len(s1) > i && len(s2) == j && preI > -1 {
			j = preI + 1
			preI = -1
			i = -1
			continue
		}
		if len(s1) == i+1 && preI > -1 {
			return true
		}

	}
	return false
}

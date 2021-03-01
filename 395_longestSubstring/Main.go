package main

import "fmt"

func main() {
	s := "a"
	k := 2
	fmt.Println(longestSubstring(s, k))
}

func longestSubstring(s string, k int) int {
	T := 0
	N := len(s)

	for i := 0; i < N; i++ {
		for step := 1; step <= N-i; step++ {
			sub := s[i : step+i]
			if isMatch(sub, k) {
				if len(sub) > T {
					T = len(sub)
				}
			}
		}
	}
	return T
}

func isMatch(s string, K int) bool {
	numMap := make(map[rune]int, 0)
	for _, e := range s {
		if _, ok := numMap[e]; !ok {
			numMap[e] = 1
			continue
		}
		numMap[e]++
	}
	for k, v := range numMap {
		if k >= 'a' && v < K {
			return false
		}
	}
	return true
}

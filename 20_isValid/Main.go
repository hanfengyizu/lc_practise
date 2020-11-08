package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}
func isValid_v2(s string) bool {
	len_origin := 1
	len_new := -len_origin
	for len_new != len_origin {
		len_origin = len(s)
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		len_new = len(s)
	}
	if len(s) != 0 {
		return false
	}

	return true
}

func isValid_v1(s string) bool {
	len_origin := len(s)
	if len_origin == 0 {
		return true
	}
	s = strings.ReplaceAll(s, "{}", "")
	s = strings.ReplaceAll(s, "()", "")
	s = strings.ReplaceAll(s, "[]", "")
	len_new := len(s)
	if len_new == len_origin {
		return false
	}
	if len_new == 0 {
		return true
	}
	return isValid(s)
}

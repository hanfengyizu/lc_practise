package main

import (
	"fmt"
	"strconv"
)

func main() {

	//"1 + 1"
	//" 2-1 + 2 "
	//"(1+(4+5+2)-3)+(6+8)"
	//"- (3 + (4 + 5))"
	fmt.Println(calculate("- (3 + (4 + 5))"))
}
func calculate(s string) int {
	sum := 0

	return sum
}
func calculate1(s string) int {
	sum := 0
	tempStr := ""
	n := len(s)
	var op byte = 0
	var op_post byte = 0
	lastIndex := 0
	j := 0
	for i := 0; i < n; {
		tempStr = ""
		lastIndex = i
		for j = i; j < n; j++ {
			if s[j] == '+' || s[j] == '-' {
				op_post = s[j]
				tempStr = tempStr + s[lastIndex:j]
				lastIndex = j + 1
				break
			}
			if s[j] == '(' || s[j] == ')' || s[j] == ' ' {
				tempStr = tempStr + s[lastIndex:j]
				lastIndex = j + 1
			}
			if j+1 == n {
				tempStr = tempStr + s[lastIndex:n]
				lastIndex = j + 1
			}
		}
		switch op {
		case '+':
			tmp, _ := strconv.Atoi(tempStr)
			sum += tmp
		case '-':
			tmp, _ := strconv.Atoi(tempStr)
			sum -= tmp
		case 0:
			tmp, _ := strconv.Atoi(tempStr)
			sum = tmp
		}
		op = op_post
		i = j + 1
	}
	return sum
}

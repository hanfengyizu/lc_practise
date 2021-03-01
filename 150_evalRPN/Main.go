package main

import (
	"fmt"
	"strconv"
)

func main() {
	test1()
	test2()
	test3()
}

func test3() {
	tokens := []string{"4", "3", "-"}
	fmt.Println(evalRPN(tokens))
}

func test2() {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}

func test1() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}
func evalRPN(tokens []string) int {
	ans := 0
	if len(tokens) == 1 {
		ans, _ = strconv.Atoi(tokens[0])
		return ans
	}
	i := 1
	A, B := 0, 0
	ii := 0
	L := len(tokens)
	for L >= 3 {
		e := ""
		A, B = 0, 0
		for ii = i - 1; ii < L; ii++ {
			e = tokens[ii]
			if e == "+" || e == "-" || e == "*" || e == "/" {
				i = ii
				A, _ = strconv.Atoi(tokens[ii-1])
				B, _ = strconv.Atoi(tokens[ii-2])
				switch e {
				case "+":
					ans = A + B
				case "-":
					ans = B - A
				case "*":
					ans = A * B
				case "/":
					ans = B / A
				}
				break
			}
		}
		tokens[i] = fmt.Sprintf("%d", ans)
		tokens = append(tokens[:i-2], tokens[i:]...)
		L = len(tokens)
	}

	return ans
}
func evalRPNV3(tokens []string) int {
	ans := 0
	if len(tokens) == 1 {
		ans, _ = strconv.Atoi(tokens[0])
		return ans
	}
	i := 1
	for len(tokens) >= 3 {
		e := ""
		A, B := 0, 0
		for ii := i - 1; ii < len(tokens); ii++ {
			e = tokens[ii]
			if e == "+" || e == "-" || e == "*" || e == "/" {
				i = ii
				A, _ = strconv.Atoi(tokens[ii-1])
				B, _ = strconv.Atoi(tokens[ii-2])
				switch e {
				case "+":
					ans = A + B
				case "-":
					ans = B - A
				case "*":
					ans = A * B
				case "/":
					ans = B / A
				}
				break
			}
		}
		tokens[i] = fmt.Sprintf("%d", ans)
		tokens = append(tokens[:i-2], tokens[i:]...)
	}

	return ans
}
func evalRPNV2(tokens []string) int {
	ans := 0
	if len(tokens) == 1 {
		ans, _ = strconv.Atoi(tokens[0])
		return ans
	}
	i := 1
	for len(tokens) >= 3 {

		e := ""
		for ii := i - 1; ii < len(tokens); ii++ {
			e = tokens[ii]
			if e == "+" || e == "-" || e == "*" || e == "/" {
				i = ii
				break
			}
		}
		A, B := 0, 0
		setA, setB := false, false
		for j := i - 1; j >= 0 && (!setA || !setB); j-- {
			if !setA {
				A, _ = strconv.Atoi(tokens[j])
				setA = true
				continue
			}
			if !setB {
				B, _ = strconv.Atoi(tokens[j])
				setB = true
				continue
			}
		}
		if i >= 2 {
			switch e {
			case "+":
				ans = A + B
			case "-":
				ans = B - A
			case "*":
				ans = A * B
			case "/":
				ans = B / A
			}
		}
		tokens[i] = fmt.Sprintf("%d", ans)
		tokens = append(tokens[:i-2], tokens[i:]...)
	}

	return ans
}
func evalRPNV1(tokens []string) int {
	xCount := 0
	ansIndex := 0
	for index, e := range tokens {
		if e == "x" {
			xCount++
			continue
		}
		ansIndex = index
	}
	ans := 0
	if xCount+1 == len(tokens) {
		ans, _ = strconv.Atoi(tokens[ansIndex])
		return ans
	}

	i := 0
	e := ""
	for i, e = range tokens {
		if e == "+" || e == "-" || e == "*" || e == "/" {
			break
		}
	}
	A, B := 0, 0
	setA, setB := false, false
	for j := i - 1; j >= 0 && (!setA || !setB); j-- {
		if tokens[j] != "x" && !setA {
			A, _ = strconv.Atoi(tokens[j])
			tokens[j] = "x"
			setA = true
			continue
		}
		if tokens[j] != "x" && !setB {
			B, _ = strconv.Atoi(tokens[j])
			tokens[j] = "x"
			setB = true
			continue
		}
	}
	if i >= 2 {
		switch e {
		case "+":
			ans = A + B
		case "-":
			ans = B - A
		case "*":
			ans = A * B
		case "/":
			ans = B / A
		}
	}
	tokens[i] = fmt.Sprintf("%d", ans)
	return evalRPN(tokens)
}

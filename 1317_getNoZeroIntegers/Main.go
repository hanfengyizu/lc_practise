package main

import "fmt"

func main() {
	fmt.Println(getNoZeroIntegers(10000))
}
func getNoZeroIntegers(n int) (ans []int) {
	A, B := 0, 0
	for i := 1; i <=(n/2 + n%2); i++ {
		A = i
		B = n - i
		if noZero(A) && noZero(B) {
			return []int{A, B}
		}
	}
	return
}

func noZero(a int) bool {
	for ; a/10 > 0; {
		if a%10 == 0 {
			return false
		}
		a = a / 10
	}
	return true
}

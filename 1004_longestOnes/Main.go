package main

import (
	"fmt"
)

//  待改进
func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
func test1() {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	fmt.Println(longestOnes(A, K))
}

func test2() {
	A := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K := 3
	fmt.Println(longestOnes(A, K))
}

func test3() {
	A := []int{0, 0, 1, 1, 1, 0, 0}
	K := 0
	fmt.Println(longestOnes(A, K))
}

func test4() {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K := 0
	fmt.Println(longestOnes(A, K))
}

func test5() {
	A := []int{1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1}
	K := 9
	fmt.Println(longestOnes(A, K))
}

func longestOnes(A []int, K int) int {
	N := len(A)
	if K == N {
		return K
	}

	if K == 0 {
		tC := 0
		tT := 0
		for _, e := range A {
			if e == 1 {
				tT++
			}
			if tT > tC {
				tC = tT
			}
			if e == 0 {
				tT = 0
			}
		}
		return tC
	}
	step := 1
	maxL := 0
	for ; step < N; step++ {
		for i := 0; i < N-step; i++ {

			count := 0
			for _, e := range A[i : i+step] {
				if e == 0 {
					count++
				}
			}
			if count <= K {
				kc := K - count
				totalOneC := step
				for j := i - 1; j >= 0; j-- {
					if A[j] == 1 {
						totalOneC++
						continue
					}
					if A[j] == 0 && kc > 1 {
						totalOneC++
						kc--
					} else {
						break
					}

				}
				for j := i + step; j < N; j++ {
					if A[j] == 1 {
						totalOneC++
						continue
					}
					if A[j] == 0 && kc > 1 {
						totalOneC++
						kc--
					} else {
						break
					}
				}
				if totalOneC > maxL {
					maxL = totalOneC
				}
			}
		}
	}
	return maxL
}

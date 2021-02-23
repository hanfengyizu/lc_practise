package main

import "fmt"

func main() {
	Test1()
	Test2()
}

func Test2() {
	customers := []int{10, 1, 7}
	grumpy := []int{0, 0, 0}
	X := 2
	fmt.Println(maxSatisfied(customers, grumpy, X))
}
func Test1() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	X := 3
	fmt.Println(maxSatisfied(customers, grumpy, X))
}
func maxSatisfied(customers []int, grumpy []int, X int) int {
	maxN := 0
	N := len(customers)
	for i := 0; i < N; i++ {
		maxN += customers[i] * (1 ^ grumpy[i])
	}
	maxS := 0
	diff := 0

	for j := 0; j < X && j < N; j++ {
		diff += customers[j] * grumpy[j]
		maxS = diff
	}

	for j := X; j < N; j++ {
		diff += customers[j]*grumpy[j] - customers[j-X]*grumpy[j-X]
		if maxS < diff {
			maxS = diff
		}
	}
	return maxS + maxN
}

func maxSatisfiedV2(customers []int, grumpy []int, X int) int {
	maxN := 0
	N := len(customers)
	for i := 0; i < N; i++ {
		if grumpy[i] == 0 {
			maxN += customers[i]
		}
	}
	maxS := 0
	diff := 0
	for i := 0; i < N; i++ {
		diff = 0
		for j := i; j < i+X && j < N; j++ {
			diff += customers[j] * grumpy[j]
		}
		if maxS < diff {
			maxS = diff
		}
	}
	return maxS + maxN
}
func maxSatisfiedV1(customers []int, grumpy []int, X int) int {
	maxS := 0
	N := len(customers)
	temp := 0
	diffSum := 0
	for i := 0; i < N; i++ {
		temp = 0
		if i <= X {
			for j := 0; j < i; j++ {
				if grumpy[j] == 0 {
					temp += customers[j]
				}
			}
			diffSum = temp
		} else {
			if grumpy[i-1] == 0 {
				diffSum += customers[i-1]
			}
			temp = diffSum
		}

		for j := i; j < i+X && j < N; j++ {
			temp += customers[j]
		}
		for j := i + X; j < N; j++ {
			if grumpy[j] == 0 {
				temp += customers[j]
			}
		}
		if maxS < temp {
			maxS = temp
		}
	}
	return maxS
}

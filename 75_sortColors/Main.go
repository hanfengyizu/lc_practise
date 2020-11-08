package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	numRed := 0
	numWhite := 0
	numBlue := 0
	for _, val := range nums {
		switch val {

		case 0:
			numRed++
			break
		case 1:
			numWhite++
			break
		case 2:
			numBlue++
			break
		}
	}
	for i := 0; i < numRed; i++ {
		nums[i] = 0
	}
	for i := numRed; i < numRed+numWhite; i++ {
		nums[i] = 1
	}
	for i := numRed + numWhite; i < numRed+numWhite+numBlue; i++ {
		nums[i] = 2
	}
}

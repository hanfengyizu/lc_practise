package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(str1))
	str1 = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(str1))
	str1 = []string{"36", "28", "70", "65", "C", "+", "33", "-46", "84", "C"}
	fmt.Println(calPoints(str1))
	str1 = []string{"8373", "C", "9034", "-17523", "D", "1492", "7558", "-5022", "C", "C", "+", "+", "-16000", "C", "+", "-18694", "C", "C", "C", "-19377", "D", "-25250", "20356", "C", "-1769", "-8303", "C", "-25332", "29884", "C", "+", "C", "-29875", "-15374", "C", "+", "14584", "13773", "+", "17037", "-28248", "+", "16822", "D", "+", "+", "-19357", "-26593", "-8548", "4776", "D", "-8244", "378", "+", "-19269", "-25447", "18922", "-16782", "2886", "C", "-17788", "D", "-22256", "C", "308", "-9185", "-19074", "1528", "28424", "D", "8658", "C", "7221", "-13704", "8995", "-21650", "22567", "-29560", "D", "-9807", "25632", "6817", "28654", "D", "-18574", "C", "D", "20103", "7875", "C", "9911", "23951", "C", "D", "C", "+", "18219", "-20922", "D", "-24923"}
	fmt.Println(calPoints(str1))
}

func calPoints(ops []string) int {
	sum := 0
	ops1 := make([]int, len(ops))
	index := -1
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op {

		case "C":
			index--
			break
		case "D":
			ops1[index+1] = ops1[index] * 2
			index++
			break
		case "+":
			ops1[index+1] = ops1[index] + ops1[index-1]
			index++
			break
		default:
			index++
			ops1[index], _ = strconv.Atoi(op)
		}
	}
	for i := 0; i <= index; i++ {
		sum += ops1[i]
	}
	return sum
}

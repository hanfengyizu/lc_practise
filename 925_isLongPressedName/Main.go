package main

import (
	"fmt"
)

func main() {
	fmt.Println(isLongPressedName("alex", "aaleeex"))
	fmt.Println(isLongPressedName("xnhtq", "xhhttqq"))
	fmt.Println(isLongPressedName("pyplrz", "ppyypllr"))
}

func isLongPressedName(name string, typed string) bool {
	nameArr := name[:]
	typedArr := typed[:]
	startIndex := 0
	endIndex := 0
	start := 0
	end := 0
	for ; start < len(nameArr); {

		for end = start + 1; end < len(nameArr) && nameArr[start] == nameArr[end]; {
			end++
		}
		if startIndex+1 > len(typedArr) || nameArr[start] != typedArr[startIndex] {
			return false
		}
		for ; startIndex < len(typedArr); {

			for endIndex = startIndex + 1; endIndex < len(typedArr) && typedArr[startIndex] == typedArr[endIndex]; {
				endIndex++
			}
			if (endIndex - startIndex) < (end - start) {
				return false
			}
			startIndex = endIndex
			break
		}
		start = end
	}
	if startIndex+1 <= len(typedArr) || start+1 <= len(nameArr) {
		return false
	}
	return true
}

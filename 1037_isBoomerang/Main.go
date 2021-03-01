package main

import "fmt"

func main() {
	arr := [][]int{{0, 1}, {2, 1}, {0, 0}}
	fmt.Println(isBoomerang(arr))
}
func isBoomerang(points [][]int) bool {
	if len(points) < 3 {
		return false
	}
	diffH := float32(points[1][0] - points[0][0])

	diffW := float32(points[1][1] - points[0][1])
	diffH1 := float32(points[2][0] - points[1][0])
	diffW1 := float32(points[2][1] - points[1][1])
	if diffH == 0 && diffW != 0 {
		if diffH1 != 0 {
			return true
		}
		return false
	}
	if diffH == 0 && diffW == 0 {
		return false
	}

	if diffH1 == 0 && diffW1 != 0 {
		return true
	}
	if diffH1 == 0 && diffW1 == 0 {
		return false
	}
	if (1.0*diffW)/(1.0*diffH) == (1.0*diffW1)/(1.0*diffH1) {
		return false
	}
	return true
}

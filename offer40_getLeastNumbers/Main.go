package main

func main() {

}
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}

	arr = append([]int{0}, arr...)
	ans := make([]int, k)
loop:
	n := len(arr)
	for i := 1; i < n; i++ {
		if 2*i < n && arr[2*i] < arr[i] {
			arr[i], arr[2*i] = arr[2*i], arr[i]
		}
		if (2*i+1) < n && arr[2*i+1] < arr[i] {
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
		}
		if 2*i < n && (2*i+1) < n && arr[2*i] >= arr[2*i+1] {
			arr[2*i], arr[2*i+1] = arr[2*i+1], arr[2*i]
		}
	}

	ans[k-1] = arr[1]
	k--
	if k > 0 {
		arr = append(arr[0:1], arr[2:]...)
		goto loop
	}
	return ans
}

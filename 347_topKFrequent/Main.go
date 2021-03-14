package main

import "fmt"

func main() {
	nums := []int{1, 2}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
func topKFrequent(nums []int, k int) []int {
	nmap := make(map[int]int)
	for _, e := range nums {
		v, ok := nmap[e]
		if ok {
			nmap[e] = v + 1
		} else {
			nmap[e] = 1
		}
	}
	nnmap := make(map[int]int)
	arr := make([]int, k+1)
	index := 0
	pushK := func(x int) {
		index++
		if index < k+1 {
			arr[index] = x
			if index < k {
				return
			}
		} else if index >= k+1 {
			// 移除当前最小元素
			if x < arr[1] {
				return
			}
			arr[1] = x
		}
		n := len(arr)
		for i := 1; i < n; {
			if 2*i < n && arr[2*i] < arr[i] {
				arr[i], arr[2*i] = arr[2*i], arr[i]
			}
			if (2*i+1) < n && arr[2*i+1] < arr[i] {
				arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
			}
			if 2*i < n && (2*i+1) < n && arr[2*i+1] < arr[2*i] {
				arr[2*i], arr[2*i+1] = arr[2*i+1], arr[2*i]
			}
			if i/2 > 0 && arr[i/2] > arr[i] {
				arr[i], arr[i/2] = arr[i/2], arr[i]
				i = i / 2
				continue
			}
			i++
		}
	}
	for k, v := range nmap {
		pushK(v)
		nnmap[v] = k
	}
	ans := make([]int, k)
	KK := 0
	for i, Key := range arr[1:] {

		for K, V := range nmap {
			if V == Key {
				ans[i] = K
				KK = K
				break
			}
		}
		delete(nmap, KK)
	}
	return ans
}

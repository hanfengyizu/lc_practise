package main

func main() {

}

type KthLargest struct {
	arr   []int
	k     int
	pushK func(x int)
	index int
}

func Constructor(k int, nums []int) KthLargest {
	ins := KthLargest{}
	ins.arr = make([]int, k+1)
	ins.pushK = func(x int) {
		ins.index++
		if ins.index < k+1 {
			ins.arr[ins.index] = x
			if ins.index < k {
				return
			}
		} else if ins.index >= k+1 {
			// 移除当前最小元素
			if x < ins.arr[1] {
				return
			}
			ins.arr[1] = x
		}
		n := len(ins.arr)
		for i := 1; i < n; {
			if 2*i < n && ins.arr[2*i] < ins.arr[i] {
				ins.arr[i], ins.arr[2*i] = ins.arr[2*i], ins.arr[i]
			}
			if (2*i+1) < n && ins.arr[2*i+1] < ins.arr[i] {
				ins.arr[i], ins.arr[2*i+1] = ins.arr[2*i+1], ins.arr[i]
			}
			if 2*i < n && (2*i+1) < n && ins.arr[2*i+1] < ins.arr[2*i] {
				ins.arr[2*i], ins.arr[2*i+1] = ins.arr[2*i+1], ins.arr[2*i]
			}
			if i/2 > 0 && ins.arr[i/2] > ins.arr[i] {
				ins.arr[i], ins.arr[i/2] = ins.arr[i/2], ins.arr[i]
				i = i / 2
				continue
			}
			i++
		}
	}
	for _, e := range nums {
		ins.pushK(e)
	}

	return ins
}

func (ins *KthLargest) Add(val int) int {
	ins.pushK(val)
	return ins.arr[1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

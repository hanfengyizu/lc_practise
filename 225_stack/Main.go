package main

import "fmt"

func main() {
	/**
	 * Your MyStack object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Push(x);
	 * param_2 := obj.Pop();
	 * param_3 := obj.Top();
	 * param_4 := obj.Empty();
	 */

	statck := MyStack{}
	statck.Push(1)
	statck.Push(2)
	fmt.Println(statck.Top())
	fmt.Println(statck.Pop())
	fmt.Println(statck.Empty())
	fmt.Println(statck.Top())
	fmt.Println(statck.Pop())
	fmt.Println(statck.Empty())

}

type MyStack struct {
	arr  []int
	size int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if len(this.arr)/(this.size+1) > 2 {
		copy(this.arr[:this.size], this.arr[:this.size])
	}
	if cap(this.arr) == 0 {
		this.arr = append(this.arr, 0)
	} else {
		tmp := this.arr[:this.size]
		this.arr = make([]int, this.size+1, this.size+1)
		copy(this.arr, tmp)
	}
	/*
		if cap(this.arr) == 0 {
			this.arr = append(this.arr, 0)
		}
		if cap(this.arr) <= this.size && cap(this.arr) > 0 {
			this.arr = append(this.arr, 0)
			//this.arr = append(this.arr, this.arr[:len(this.arr)/2+1]...)
		}*/

	this.arr[this.size] = x
	this.size++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	this.size--
	return this.arr[this.size]
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.arr[this.size-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.size <= 0
}

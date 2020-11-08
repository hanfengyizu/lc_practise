package main

import "fmt"

func main() {
	fmt.Println(2147483647 <= 2147483648)
	fmt.Println("*****************")
	main4()
	fmt.Println("*****************")
	main3()
	fmt.Println("*****************")
	main2()
	fmt.Println("*****************")
	main1()
}
func main4() {
	minStack := MinStack{}
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Push(-2147483648)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
func main3() {
	minStack := MinStack{}
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
func main2() {
	minStack := MinStack{}
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
func main1() {
	minStack := MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	origin []int
	statck []int

	statck_len int
	origin_len int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.origin = append(this.origin, 0)
	this.statck = append(this.statck, 0)
	this.origin[this.origin_len] = x
	this.origin_len = this.origin_len + 1
	this.statck[this.statck_len] = 0
	this.statck_len = this.statck_len + 1
	if this.statck_len == 1 {
		this.statck[0] = x
		return
	}
	for i := this.statck_len - 2; i >= 0; i-- {
		if x <= this.statck[i] {
			this.statck[i+1] = x
			return
		}
		this.statck[i+1] = this.statck[i]
	}

	this.statck[0] = x
}

func (this *MinStack) Pop() {
	deleEle := this.origin[this.origin_len-1]
	this.origin_len--
	for i := this.statck_len - 1; i >= 0; i-- {
		if deleEle == this.statck[i] {
			dele_index := i
			for ; dele_index < this.statck_len-1; dele_index++ {
				this.statck[dele_index] = this.statck[dele_index+1]
			}
			this.statck_len--
			break
		}
	}

}

func (this *MinStack) Top() int {
	return this.origin[this.origin_len-1]
}

func (this *MinStack) GetMin() int {
	return this.statck[this.statck_len-1]
}

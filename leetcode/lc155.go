package leetcode

import "math"

type MinStack struct {
	min   int
	stack []int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		min:   math.MaxInt,
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if pop == this.min {
		this.min = math.MaxInt
		for _, v := range this.stack {
			if v < this.min {
				this.min = v
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

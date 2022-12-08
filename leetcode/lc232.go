package leetcode

type MyQueue struct {
	queue  []int
	length int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		queue:  make([]int, 0),
		length: 0,
	}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
	this.length++
}

func (this *MyQueue) Pop() int {
	pop := this.queue[0]
	this.queue = this.queue[1:]
	this.length--
	return pop
}

func (this *MyQueue) Peek() int {
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return this.length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

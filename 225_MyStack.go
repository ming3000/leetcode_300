package leetcode_300

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func MyStackConstructor() MyStack {
	return MyStack{queue: make([]int, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	// 新元素入队
	// 除了新元素之外，其余元素都从队首出队后再从队尾入队，即完成了反序
	this.queue = append(this.queue, x)
	length := len(this.queue)
	for i := 0; i < length-1; i++ {
		head := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, head)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ret := this.queue[0]
	this.queue = this.queue[1:]
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

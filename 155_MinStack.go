package leetcode_300

type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	minDataLen := len(this.minData)
	if minDataLen == 0 || this.minData[minDataLen-1] > x {
		this.minData = append(this.minData, x)
	} else {
		this.minData = append(this.minData, this.minData[minDataLen-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
	this.minData = this.minData[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
}

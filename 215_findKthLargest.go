package leetcode_300

import "container/heap"

type BigTopArray []int

func (b BigTopArray) Len() int {
	return len(b)
}

func (b BigTopArray) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BigTopArray) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *BigTopArray) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *BigTopArray) Pop() interface{} {
	ret := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return ret
}

func findKthLargest(nums []int, k int) int {
	temp := make(BigTopArray, 0, k)
	var count int
	for _, v := range nums {
		if count < k {
			heap.Push(&temp, v)
			count++
		} else {
			if v > temp[0] {
				heap.Pop(&temp)
				heap.Push(&temp, v)
			} //>>>
		} //>>
	} //>

	return temp[0]
}

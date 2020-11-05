package leetcode_300

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// make a cycle and count the length of list
	var length = 1
	cur := head
	for cur.Next != nil {
		length++
		cur = cur.Next
	}
	cur.Next = head

	cur = head
	for i := 0; i < length-k%length-1; i++ {
		cur = cur.Next
	} //>

	newHead := cur.Next
	cur.Next = nil
	return newHead
}

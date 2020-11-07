package mark

func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 奇数个
	if fast != nil {
		slow = slow.Next
	}
	var pre, cur *ListNode
	cur = slow
	for cur != nil {
		nextBack := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextBack
	} //>

	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		} //>>
		head = head.Next
		pre = pre.Next
	} //>

	return true
}

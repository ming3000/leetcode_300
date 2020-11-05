package leetcode_300

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		} //>>
		slow = slow.Next
		fast = fast.Next.Next
	} //>
	return false
}

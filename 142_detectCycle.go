package leetcode_300

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			for head != fast {
				head = head.Next
				fast = fast.Next
			} //>>>
			return head
		} //>>
		slow = slow.Next
		fast = fast.Next.Next
	} //>

	return nil
}

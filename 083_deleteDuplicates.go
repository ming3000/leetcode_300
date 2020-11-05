package leetcode_300

func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	sentry := &ListNode{Next: head}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		} //>>
	} //>
	return sentry.Next
}

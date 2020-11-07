package mark

func deleteDuplicates082(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	pre, cur := sentry, head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		} //>>
		if pre.Next != cur {
			pre.Next = cur.Next
		} else {
			pre = cur
		} //>>
		cur = cur.Next
	} //>
	return sentry.Next
}

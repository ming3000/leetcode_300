package mark

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	sentry := &ListNode{}
	return mergeList(sentry, left, right)
}

func mergeList(sentry *ListNode, left, right *ListNode) *ListNode {
	cur := sentry
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		} //>>
		cur = cur.Next
	} //>
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return sentry.Next
}

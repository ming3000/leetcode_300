package mark

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	sentry := &ListNode{Next: head}
	preM := sentry
	for i := 1; i <= m-1; i++ {
		preM = preM.Next
	}

	var pre, cur *ListNode
	cur = preM.Next
	for i := m; i <= n; i++ {
		nextBack := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextBack
	}

	preM.Next.Next = cur
	preM.Next = pre
	return sentry.Next
}

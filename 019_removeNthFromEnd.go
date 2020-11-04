package leetcode_300

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{Next: head}
	var pre *ListNode
	var cur = sentry
	var count = 1

	for head != nil {
		if count >= n {
			pre = cur
			cur = cur.Next
		} //>>
		head = head.Next
		count++
	} //>
	pre.Next = cur.Next
	return sentry.Next
}

package leetcode_300

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var cur *ListNode
	var carry int
	sentry := &ListNode{}
	cur = sentry
	for l1 != nil || l2 != nil || carry != 0 {
		curSum := carry
		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		} //>>
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		} //>>
		cur.Next = &ListNode{Val: curSum % 10}
		cur = cur.Next
		carry = curSum / 10
	} //>

	return sentry.Next
}

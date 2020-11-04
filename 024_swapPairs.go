package leetcode_300

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentry := &ListNode{Next: head}
	pre := sentry
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		first, second := pre.Next, pre.Next.Next
		first.Next = second.Next
		second.Next = first
		pre.Next = second

		pre = first
	} //>
	return sentry.Next
}

package leetcode_300

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		nextBack := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextBack
	}
	return pre
}

package leetcode_300

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
	l := sortList(head)
	r := sortList(mid)
	root := &ListNode{}
	return mergeList(root, l, r)
}

func mergeList(head, left, right *ListNode) *ListNode {
	tail := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		} //>>
		tail = tail.Next
	} //>

	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}
	return head.Next
}

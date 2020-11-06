package leetcode_300

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha, hb := headA, headB
	var la, lb int
	for ha != nil {
		la++
		ha = ha.Next
	}
	for hb != nil {
		lb++
		hb = hb.Next
	}

	var longHead, shortHead *ListNode
	var diff int
	if la > lb {
		longHead, shortHead = headA, headB
		diff = la - lb
	} else {
		longHead, shortHead = headB, headA
		diff = lb - la
	}

	for longHead != nil {
		if diff > 0 {
			longHead = longHead.Next
			diff--
			continue
		} //>>
		if longHead != shortHead {
			longHead, shortHead = longHead.Next, shortHead.Next
		} else {
			break
		} //>>
	} //>
	return longHead
}

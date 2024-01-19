func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		left  = make(map[*ListNode]struct{})
	)

	for headA != nil {
		left[headA]= struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		_, ok := left[headB]
		if ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
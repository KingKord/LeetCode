func removeNthFromEnd(head *ListNode, n int) *ListNode {
	 if rec(head, n) == 0 {
		 return head.Next
	 }
	return head
}

func rec(head *ListNode, left int) int {
	if head == nil {
		return left
	}

	left = rec(head.Next, left)
	if left == 0 {
		head.Next = head.Next.Next
	}
	return left - 1
}
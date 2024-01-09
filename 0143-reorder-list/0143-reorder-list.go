func recursion(head, now *ListNode) *ListNode {
	if now == nil {
		return head
	}

	left := recursion(head, now.Next)

	if left == nil {
		return nil
	} else if left == now {
		left.Next = nil
		return nil
	} else if left.Next == now {
		left.Next.Next = nil
		return nil
	}

	temp := left.Next
	left.Next = now
	left.Next.Next = temp

	return left.Next.Next
}

func reorderList(head *ListNode) {
	recursion(head, head)
}
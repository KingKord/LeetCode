
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		heap []int
		ans  = &ListNode{
			Val:  math.MinInt,
			Next: nil,
		}
		length = 0
		temp   = ans
	)

	for _, node := range lists {
		for node != nil {
			heap = appendToHeap(heap, node.Val)
			node = node.Next
		}
	}
	length = len(heap)
	for i := 0; i < length; i++ {
		var value int
		heap, value = retrieveFromHeap(heap)
		temp.Next = &ListNode{Val: value}
		temp = temp.Next
	}

	return ans.Next
}

func retrieveFromHeap(heap []int) ([]int, int) {
	last := len(heap) - 1
	val := heap[0]
	heap[0], heap[last] = heap[last], heap[0]
	heap = heap[:last]
	idx, left, right := 0, 1, 2
	for {
		if left > len(heap)-1 {
			break
		} else if left == len(heap)-1 {
			if heap[idx] > heap[left] {
				heap[idx], heap[left] = heap[left], heap[idx]
			}
			break
		}

		if heap[left] < heap[right] {
			if heap[idx] <= heap[left] {
				break
			}
			heap[idx], heap[left] = heap[left], heap[idx]
			idx = left
		} else {
			if heap[idx] <= heap[right] {
				break
			}

			heap[idx], heap[right] = heap[right], heap[idx]
			idx = right
		}
		left = 2*idx + 1
		right = 2*idx + 2
	}

	return heap, val
}

func appendToHeap(heap []int, val int) []int {
	heap = append(heap, val)
	var (
		child  = len(heap) - 1
		parent = (child - 1) / 2
	)

	for parent != -1 && heap[parent] > heap[child] {
		heap[parent], heap[child] = heap[child], heap[parent]

		child = parent
		parent = (child - 1) / 2
	}
	return heap
}

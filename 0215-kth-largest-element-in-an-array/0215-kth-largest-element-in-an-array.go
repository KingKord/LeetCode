func findKthLargest(nums []int, k int) int {
	var heap []int
	heap = append(heap, nums[0])

	for i := 1; i < len(nums); i++ {
		value := nums[i]
		if len(heap) < k {
			heap = appendToHeap(heap, value)
		} else if nums[i] > heap[0] {
			heap = retrieveFromHeap(heap, value)
		}
	}

	return heap[0]
}

func retrieveFromHeap(heap []int, value int) []int {
	heap[0] = value

	idx, leftChild, rightChild := 0, 1, 2
	for {
		if leftChild > len(heap)-1 {
			break
		}
		if leftChild == len(heap)-1 {
			if heap[leftChild] < heap[idx] {
				heap[leftChild], heap[idx] = heap[idx], heap[leftChild]
			}
			break
		}

		if heap[leftChild] < heap[rightChild] {
			if heap[leftChild] >= heap[idx] {
				break
			}

			heap[leftChild], heap[idx] = heap[idx], heap[leftChild]
			idx = leftChild
		} else {
			if heap[rightChild] >= heap[idx] {
				break
			}

			heap[rightChild], heap[idx] = heap[idx], heap[rightChild]
			idx = rightChild
		}

		leftChild = 2*idx + 1
		rightChild = 2*idx + 2
	}

	return heap
}

func appendToHeap(heap []int, value int) []int {
	heap = append(heap, value)
	idx := len(heap) - 1
	parent := (idx - 1) / 2

	for heap[parent] > heap[idx] {
		heap[parent], heap[idx] = heap[idx], heap[parent]

		idx = parent
		parent = (idx - 1) / 2
	}
	return heap
}
func lastStoneWeight(stones []int) int {
	var heap []int
	for _, stone := range stones {
		heap = appendToHeap(heap, stone)
	}

	var y, x int

	for len(heap) > 1 {
		heap, y = retrieveFromHeap(heap)
		heap, x = retrieveFromHeap(heap)

		if x != y {
			heap = appendToHeap(heap, y-x)
		}
	}
	if len(heap) == 0 {
		return 0
	}
	return heap[0]
}

func retrieveFromHeap(heap []int) ([]int, int) {
	value := heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	idx, leftChild, rightChild := 0, 1, 2
	for {
		if leftChild > len(heap)-1 {
			break
		}
		if leftChild == len(heap)-1 {
			if heap[leftChild] > heap[idx] {
				heap[leftChild], heap[idx] = heap[idx], heap[leftChild]
			}
			break
		}

		if heap[leftChild] > heap[rightChild] {
			if heap[leftChild] <= heap[idx] {
				break
			}

			heap[leftChild], heap[idx] = heap[idx], heap[leftChild]
			idx = leftChild
		} else {
			if heap[rightChild] <= heap[idx] {
				break
			}

			heap[rightChild], heap[idx] = heap[idx], heap[rightChild]
			idx = rightChild
		}

		leftChild = 2*idx + 1
		rightChild = 2*idx + 2
	}

	return heap, value
}

func appendToHeap(heap []int, value int) []int {
	heap = append(heap, value)
	idx := len(heap) - 1
	parent := (idx - 1) / 2

	for heap[parent] < heap[idx] {
		heap[parent], heap[idx] = heap[idx], heap[parent]

		idx = parent
		parent = (idx - 1) / 2
	}
	return heap
}

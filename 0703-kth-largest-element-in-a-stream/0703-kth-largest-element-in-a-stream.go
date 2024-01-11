type KthLargest struct {
	heap []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k: k,
	}
	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) == 0 {
		this.heap = append(this.heap, val)
		return val
	}

	if len(this.heap) < this.k {
		this.heap = appendToHeap(this.heap, val)
		return this.heap[0]
	}
	if val <= this.heap[0] {
		return this.heap[0]
	}

	this.heap = down(this.heap, val)
	return this.heap[0]
}

func down(heap []int, val int) []int {
	heap[0] = val
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

	return heap
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

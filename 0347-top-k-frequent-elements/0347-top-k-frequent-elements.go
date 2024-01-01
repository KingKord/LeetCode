func topKFrequent(nums []int, k int) []int {
	var (
		heap       []int
		idxToValue = make(map[int]int)
		valueToIdx = make(map[int]int)
	)

	for _, num := range nums {
		idx, ok := valueToIdx[num]
		if !ok {
			heap = append(heap, 1)
			idxToValue[len(heap)-1] = num
			valueToIdx[num] = len(heap) - 1
		} else {
			heap[idx]++

			parentIdx := (idx - 1) / 2
			for heap[parentIdx] < heap[idx] {
				swap(heap, idxToValue, valueToIdx, parentIdx, idx)

				idx = parentIdx
				parentIdx = (idx - 1) / 2
			}
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = idxToValue[0]

		swap(heap, idxToValue, valueToIdx, 0, len(heap)-1)
		heap = heap[:len(heap)-1]
		idx, left, right := 0, 1, 2

		for {

			if left > len(heap)-1 {
				break
			} else if left == len(heap)-1 {
				max := heap[left]
				if heap[idx] < max {
					swap(heap, idxToValue, valueToIdx, idx, left)
				}
				break
			}
			if heap[left] > heap[right] {

				max := heap[left]
				if heap[idx] >= max {
					break
				}

				swap(heap, idxToValue, valueToIdx, idx, left)
				idx = left
			} else {
				max := heap[right]
				if heap[idx] >= max {
					break
				}

				swap(heap, idxToValue, valueToIdx, idx, right)
				idx = right
			}

			left = 2*idx + 1
			right = 2*idx + 2
		}
	}
	return ans
}

func swap(heap []int, idxToValue map[int]int, valueToIdx map[int]int, parentIdx int, idx int) {
	heap[idx], heap[parentIdx] = heap[parentIdx], heap[idx]

	child, parent := idxToValue[idx], idxToValue[parentIdx]
	valueToIdx[child], valueToIdx[parent] = valueToIdx[parent], valueToIdx[child]

	idxToValue[idx], idxToValue[parentIdx] = parent, child
}
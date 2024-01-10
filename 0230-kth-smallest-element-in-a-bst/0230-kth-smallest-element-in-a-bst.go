func kthSmallest(root *TreeNode, k int) int {
	_, val := rec(root, k, 0)
	return val
}

func rec(root *TreeNode, k, n int) (int, int) {
	if root == nil {
		return n, 0
	}

	leftCount, leftValue := rec(root.Left, k, n)
	if leftCount == k {
		return leftCount, leftValue
	}
	leftCount++
	if leftCount == k {
		return leftCount, root.Val
	}
	rightCount, rightValue := rec(root.Right, k, leftCount)
	if rightCount == k {
		return rightCount, rightValue
	}

	return rightCount, 0
}
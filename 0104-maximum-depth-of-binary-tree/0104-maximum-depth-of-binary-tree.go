func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depthLeft, depthRight int

	depthLeft = maxDepth(root.Left) + 1
	depthRight = maxDepth(root.Right) + 1

    return Max(depthLeft, depthRight)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

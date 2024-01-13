func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	
	var depth = 0
	for _, child := range root.Children {
		depth = max(depth, maxDepth(child))
	}
	return depth + 1
}
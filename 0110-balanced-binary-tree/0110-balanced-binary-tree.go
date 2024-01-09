func rec(root *TreeNode) (int,bool) {
	if root == nil {
		return 0,true
	}

	leftDepth,leftBalanced := rec(root.Left)
	rightDepth,rightBalanced := rec(root.Right)
	sub := leftDepth - rightDepth
	if !leftBalanced || !rightBalanced {
		return 0, false
	}
	return max(leftDepth, rightDepth) + 1, -1 <= sub && sub <= 1
}
func isBalanced(root *TreeNode) bool {
	_, balanced := rec(root)
	return balanced
}

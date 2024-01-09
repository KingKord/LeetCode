func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	var result bool
	if root.Val == subRoot.Val {
		result = check(root, subRoot)
	}

	isLeft := isSubtree(root.Left, subRoot)
	isRight := isSubtree(root.Right, subRoot)

	return isRight || isLeft || result
}

func check(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}
	left := check(root.Left, subRoot.Left)
	right := check(root.Right, subRoot.Right)

	return right && left
}

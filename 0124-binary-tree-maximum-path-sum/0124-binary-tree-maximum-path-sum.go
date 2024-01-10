func maxPathSum(root *TreeNode) int {
	_, m := recursion(root)
	return m
}

func recursion(root *TreeNode) (int, int) {
	var leftNow, rightNow, leftMax, rightMax, now int
	now = root.Val
	if root.Left != nil {
		leftNow, leftMax = recursion(root.Left)
		if leftNow > 0 {
			now += leftNow
		}
	} else {
		leftMax = math.MinInt
	}
	if root.Right != nil {
		rightNow, rightMax = recursion(root.Right)
		if rightNow > 0 {
			now += rightNow
		}
	} else {
		rightMax = math.MinInt
	}

	maxLast := max(leftMax, rightMax, now)
	if root.Left != nil && root.Right != nil {
		now = max(root.Val+leftNow, root.Val+rightNow, root.Val)
	} else if root.Left == nil && root.Right != nil {
		now = max(root.Val+rightNow, root.Val)
	} else 	if root.Left != nil && root.Right == nil {
		now = max(root.Val+leftNow, root.Val)
	} else {
		now = root.Val
	}
	return now, maxLast
}
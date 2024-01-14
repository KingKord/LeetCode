func recursion(root *TreeNode, i int, ans [][]int) [][]int {
	if root == nil {
		return ans
	}
	
	if len(ans) <= i {
		ans = append(ans, []int{root.Val})
	} else {
		ans[i] = append(ans[i], root.Val)
	}

	ans = recursion(root.Left, i+1, ans)
	ans = recursion(root.Right, i+1, ans)

	return ans
}

func levelOrder(root *TreeNode) [][]int {
	return recursion(root, 0, nil)
}
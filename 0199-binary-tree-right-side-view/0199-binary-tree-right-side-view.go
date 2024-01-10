func rightSideView(root *TreeNode) []int {
	return rec(root, 0, []int{})
}

func rec(root *TreeNode, depth int, ans []int) []int {
	if root == nil {
		return ans
	}
	if depth == len(ans) {
		ans = append(ans, root.Val)
	}
    
	ans = rec(root.Right, depth+1, ans)
	ans = rec(root.Left, depth+1, ans)
	return ans
}
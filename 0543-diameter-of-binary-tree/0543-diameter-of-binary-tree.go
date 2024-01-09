func rec(root *TreeNode) (int,int) {
	if root == nil {
		return 0, 0
	}
	maxLeft, depthLeft := rec(root.Left)
	maxRight, depthRight := rec(root.Right)

    return  max(maxLeft,maxRight, depthLeft + depthRight), max(depthLeft,depthRight) +1
}
func diameterOfBinaryTree(root *TreeNode) int {
	max, _ := rec(root)
	return max
}
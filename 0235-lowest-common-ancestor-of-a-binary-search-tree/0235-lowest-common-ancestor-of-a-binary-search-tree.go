func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	left := lowestCommonAncestor(root.Left, p,q)
	right := lowestCommonAncestor(root.Right, p,q)
	
	if left != nil && right != nil {
		return root
	}
	
	if (root.Val == p.Val || root.Val == q.Val) && (left != nil || right != nil) {
		return root
	}
	
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	
	return nil
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	
	if p == nil || q == nil {
		return false
	}
	
	if q.Val != p.Val {
		return false
	}
 	if !isSameTree(p.Left, q.Left) || !isSameTree(p.Right,q.Right) {
		 return false
	}
 
	return true
}
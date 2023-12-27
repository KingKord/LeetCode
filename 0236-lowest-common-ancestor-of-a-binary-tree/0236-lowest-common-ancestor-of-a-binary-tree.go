func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res, _ := find(root, p, q, nil)
	fmt.Println("----------------------------------")
	return res
}

func find(root, p, q, ans *TreeNode) (*TreeNode, bool) {
	if ans != nil {
		return ans, true
	}
	var isCurrent bool
	var isFoundedLeft bool
	var isFoundedRight bool

	if root.Val == p.Val || root.Val == q.Val {
		isCurrent = true
	}

	if root.Left != nil {
		ans, isFoundedLeft = find(root.Left, p, q, ans)
		if ans != nil {
			return ans, true
		}
		if isCurrent && isFoundedLeft {
			return root, true
		}
	}

	if root.Right != nil {
		ans, isFoundedRight = find(root.Right, p, q, ans)
		if ans != nil {
			return ans, true
		}
		if isCurrent && isFoundedRight {
			return root, true
		}

		if isFoundedLeft && isFoundedRight {
			return root, true
		}
	}

	if isCurrent || isFoundedRight || isFoundedLeft {
		return nil, true
	}
	return nil, false
}

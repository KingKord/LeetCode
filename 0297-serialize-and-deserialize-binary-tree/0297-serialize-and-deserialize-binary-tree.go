
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	return strings.Join(serializeRecursion(root, []string{}), ",")
}

func serializeRecursion(root *TreeNode, ans []string) []string {
	if root == nil {
		ans = append(ans, "null")
		return ans
	}

	ans = append(ans, strconv.Itoa(root.Val))

	ans = serializeRecursion(root.Left, ans)
	ans = serializeRecursion(root.Right, ans)
	
	return ans
}

func (this *Codec) deserialize(data string) *TreeNode {
	roots := strings.Split(data, ",")
	var val int
    now := &val
	return deserializeRecursion(roots, now)
}

func deserializeRecursion(roots []string, now *int) *TreeNode {
	if roots[*now] == "null" {
		*now++
		return nil
	}
	val, _ := strconv.Atoi(roots[*now]) 
	*now++
	root := &TreeNode{
		Val:   val,
		Left:  deserializeRecursion(roots, now),
		Right: deserializeRecursion(roots, now),
	}
	
	return root
}
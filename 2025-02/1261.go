// Runtime : 3 ms
// Memory  : 8.93 MB
package feb25

type FindElements struct {
	Cache map[int]bool
}

func Constructor3(root *TreeNode) FindElements {
	root.Val = 0
	cache := make(map[int]bool)
	cache[0] = true

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
			cache[node.Left.Val] = true
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
			cache[node.Right.Val] = true
			queue = append(queue, node.Right)
		}
	}

	return FindElements{cache}
}

func (this *FindElements) Find(target int) bool {
	if seen, ok := this.Cache[target]; ok {
		return seen
	}
	return false
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

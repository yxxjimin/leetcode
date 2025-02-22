// Runtime : 0 ms
// Memory  : 6.33 MB
package feb25

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	stack, depth := []*TreeNode{}, 0
	var root *TreeNode

	i, n := 0, len(traversal)

	for i < n {
		val := 0
		for i < n && traversal[i] != '-' {
			val = val*10 + int(traversal[i]-'0')
			i++
		}

		stack = stack[:depth]
		child := &TreeNode{val, nil, nil}

		if len(stack) > 0 {
			parent := stack[len(stack)-1]
			if parent.Left == nil {
				parent.Left = child
			} else {
				parent.Right = child
			}
		} else {
			root = child
		}
		stack = append(stack, child)

		depth = 0
		for i < n && traversal[i] == '-' {
			depth++
			i++
		}
	}

	return root
}

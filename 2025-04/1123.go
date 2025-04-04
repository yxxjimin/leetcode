// Runtime : 1 ms
// Memory  : 6.83 MB
package apr25

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	leftDepth, leftLcs := dfs(root.Left)
	rightDepth, rightLcs := dfs(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1, leftLcs
	} else if leftDepth < rightDepth {
		return rightDepth + 1, rightLcs
	}
	return leftDepth + 1, root
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := dfs(root)
	return lca
}

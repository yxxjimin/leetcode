// Runtime : 0 ms
// Memory  : 4.89 MB
package feb25

func constructTree(preorder, postValToIdx []int, preStart, preEnd, postStart int) *TreeNode {
	if preStart > preEnd {
		return nil
	} else if preStart == preEnd {
		return &TreeNode{preorder[preStart], nil, nil}
	}

	leftVal := preorder[preStart+1]
	leftTreeSize := postValToIdx[leftVal] - postStart + 1

	leftTree := constructTree(preorder, postValToIdx, preStart+1, preStart+leftTreeSize, postStart)
	rightTree := constructTree(preorder, postValToIdx, preStart+leftTreeSize+1, preEnd, postStart+leftTreeSize)

	return &TreeNode{preorder[preStart], leftTree, rightTree}
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	postValToIdx := make([]int, n+1)
	for i := range postorder {
		postValToIdx[postorder[i]] = i
	}

	return constructTree(preorder, postValToIdx, 0, n-1, 0)
}

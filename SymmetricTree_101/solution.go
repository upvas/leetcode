package SymmetricTree_101

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSym(root.Left, root.Right)
}

func isSym(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == nil && node2 == nil
	} else if node1.Val != node2.Val {
		return false
	}

	return isSym(node1.Left, node2.Right) && isSym(node1.Right, node2.Left)
}
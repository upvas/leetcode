package MaximumDepthOfBinaryTree_104

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return 1 + lDepth
	} else {
		return 1 + rDepth
	}
}
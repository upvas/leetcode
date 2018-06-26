package SameTree_100

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

// BFS
func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}

	for len(queue) > 0 {
		first := queue[0]
		second := queue[1]
		queue = queue[2:]

		if first == nil && second == nil {
			continue
		}

		if first == nil || second == nil {
			return false
		}

		if first.Val != second.Val {
			return false
		}

		queue = append(queue, first.Left, second.Left, first.Right, second.Right)
	}

	return true
}
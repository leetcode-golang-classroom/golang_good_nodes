package sol

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	max := math.MinInt32
	CountGoodNodes(root, max, &count)
	return count
}

func CountGoodNodes(root *TreeNode, max int, count *int) {
	if root == nil {
		return
	}
	curMax := max
	if curMax <= root.Val {
		*count += 1
		curMax = root.Val
	}
	CountGoodNodes(root.Left, curMax, count)
	CountGoodNodes(root.Right, curMax, count)
}

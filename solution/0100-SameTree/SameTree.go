package leetcode

import "LeetCode-Go/utils"

func isSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

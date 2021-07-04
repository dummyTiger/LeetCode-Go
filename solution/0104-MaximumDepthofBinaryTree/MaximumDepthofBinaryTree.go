package leetcode
import "LeetCode-Go/utils"

func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left) + 1
	maxRight := maxDepth(root.Right) + 1
	if maxLeft > maxRight {
		return maxLeft
	} else {
		return maxRight
	}

}

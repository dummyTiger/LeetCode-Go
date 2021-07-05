package leetcode

import "LeetCode-Go/utils"

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	preNode := head
	node := head.Next
	for {
		if node.Val == preNode.Val {
			preNode.Next = node.Next
		} else {
			preNode = node
		}
		node = node.Next
		if node == nil {
			break
		}
	}
	return head
}

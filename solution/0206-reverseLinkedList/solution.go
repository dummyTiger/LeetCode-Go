package _206_reverseLinkedList

import "LeetCode-Go/utils"

func reverseList(head *utils.ListNode) *utils.ListNode {
	cur := head
	var pre *utils.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

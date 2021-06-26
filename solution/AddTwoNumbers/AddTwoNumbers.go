package leetcode

import "github.com/dummyTiger/LeetCode-Go/utils"

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	head := &utils.ListNode{Val: 0}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node := &utils.ListNode{Val: sum % 10}
		current.Next = node
		current = current.Next
		carry = sum / 10
	}
	return head.Next
}

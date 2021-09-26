package leetcode

import "LeetCode-Go/utils"

func hasCycle(head *utils.ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow := head.Next
	quick := head.Next.Next

	for ; slow != nil && quick != nil; {

		if slow == quick {
			return true
		}
		slow = slow.Next
		if quick.Next == nil {
			return false
		}
		quick = quick.Next.Next
	}

	return false
}

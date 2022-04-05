package _142_LinkedListCycleII

import "LeetCode-Go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := slow

	for ; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			index1 := fast
			index2 := head
			for ; index1 != index2; {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}

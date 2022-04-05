package _019_RemoveNthNodeFromEndofList

import "LeetCode-Go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {

	dummy := &utils.ListNode{Val: 0}

	dummy.Next = head

	front := head
	back := dummy

	for i := 0; i < n; i++ {
		front = front.Next
	}

	for front != nil {
		front = front.Next
		back = back.Next
	}

	back.Next = back.Next.Next
	return dummy.Next
}

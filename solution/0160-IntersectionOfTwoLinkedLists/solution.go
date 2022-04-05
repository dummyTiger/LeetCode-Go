package _160_IntersectionOfTwoLinkedLists

import "LeetCode-Go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	cA, cB := headA, headB
	for cA != nil && cB != nil {
		cA, cB = cA.Next, cB.Next
	}

	var kA, kB int
	for cA != nil {
		cA = cA.Next
		kA++
	}
	for cB != nil {
		cB = cB.Next
		kB++
	}

	if kA > 0 {
		for i := 0; i < kA; i++ {
			headA = headA.Next
		}
	}
	if kB > 0 {
		for i := 0; i < kB; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA, headB = headA.Next, headB.Next
	}
	return headA
}

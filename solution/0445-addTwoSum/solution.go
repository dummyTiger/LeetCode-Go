package _445_addTwoSum

import "LeetCode-Go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	n1 := reverseList(l1)
	n2 := reverseList(l2)

	dummy := &utils.ListNode{}
	sumNode := dummy
	carry := 0
	for n1 != nil || n2 != nil {
		sum := 0 + carry
		if n1 != nil {
			sum += n1.Val
		}
		if n2 != nil {
			sum += n2.Val
		}
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		newNode := &utils.ListNode{Val: sum}
		sumNode.Next = newNode
		sumNode = sumNode.Next
		if n1 != nil {
			n1 = n1.Next
		}

		if n2 != nil {
			n2 = n2.Next
		}
	}

	if carry > 0 {
		sumNode.Next = &utils.ListNode{Val: carry}
	}

	return reverseList(dummy.Next)
}

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

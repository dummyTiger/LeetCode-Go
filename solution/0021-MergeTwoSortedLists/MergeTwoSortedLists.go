package leetcode

import "LeetCode-Go/utils"

func mergeTwoLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {

	if l1 == nil && l2 == nil {
		return nil

	}

	if l1 == nil && l2!=nil {
		return l2
	}

	if l1 !=nil && l2 ==nil {
		return l1
	}

	values:=make([]int,0)


	for ;l2!=nil && l1!=nil; {
		if l1.Val<l2.Val {
			values = append(values,l1.Val)
			l1 = l1.Next
		}else {
			values = append(values,l2.Val)
			l2 = l2.Next
		}
	}

	if l1 !=nil {
		for;l1!=nil; {
			values = append(values,l1.Val)
			l1 = l1.Next
		}

	}

	if l2 !=nil {
		for;l2!=nil; {
			values = append(values,l2.Val)
			l2 = l2.Next
		}
	}

	head1 := &utils.ListNode{Val: 0}
	current := head1
	for idx, i := range values {

		if idx == 0 {
			head1.Val = i
			continue
		}
		current.Next = &utils.ListNode{Val: i}
		current = current.Next
	}

	return head1


}

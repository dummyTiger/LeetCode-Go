package leetcode

import (
	"LeetCode-Go/utils"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{9, 9, 9, 9, 9, 9, 9}

	head1 := &utils.ListNode{Val: 0}
	current := head1
	head2 := &utils.ListNode{Val: 0}
	for idx, i := range array1 {

		if idx == 0 {
			head1.Val = i
			continue
		}
		current.Next = &utils.ListNode{Val: i}
		current = current.Next
	}

	for idx, i := range array2 {

		if idx == 0 {
			head2.Val = i
			continue
		}
		current.Next = &utils.ListNode{Val: i}
		current = current.Next
	}

	numbers := addTwoNumbers(head1, head2)

	result := []int{0, 2, 3, 4, 5, 0, 0, 1}

	current = numbers

	for i := 0; current != nil && i < len(result); {
		if current.Val != result[i] {
			t.Errorf("result not eaquel")
		}
		current = current.Next
	}
}

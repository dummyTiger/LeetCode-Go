package leetcode

import (
	"LeetCode-Go/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
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
	current = head2

	for idx, i := range array2 {

		if idx == 0 {
			head2.Val = i
			continue
		}
		current.Next = &utils.ListNode{Val: i}
		current = current.Next
	}

	h := mergeTwoLists(head1, head2)

	assert.Equal(t,h.Val,1)
}

package _445_addTwoSum

import (
	"LeetCode-Go/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}

	addTwoNumbers(&utils.ListNode{Val: 9, Next: nil}, &utils.ListNode{Val: 1, Next: nil})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addTwoNumbers(&utils.ListNode{Val: 0, Next: nil}, &utils.ListNode{Val: 0, Next: nil})
		})
	}
}

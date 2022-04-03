package _167_TwoSumII

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		inputs []int
		target int
		expect []int
	}{
		{inputs: []int{2, 7, 11, 15}, target: 9, expect: []int{1, 2}},
		{inputs: []int{2, 3, 4}, target: 6, expect: []int{1, 3}},
	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCases {
			assert.Equal(t, test.expect, twoSum(test.inputs, test.target))
		}
	})
}

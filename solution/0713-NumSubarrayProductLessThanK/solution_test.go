package _713_NumSubarrayProductLessThanK

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		inputs []int
		target int
		expect int
	}{
		{inputs: []int{10, 5, 2, 6}, target: 100, expect: 8},
		{inputs: []int{1, 2, 3}, target: 0, expect: 0},
	}

	for _, test := range testCases {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, test.expect, numSubarrayProductLessThanK(test.inputs, test.target))
		})
	}
}

package _560_SubarraySum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	testCase := []struct {
		input  []int
		target int
		expect int
	}{
		{input: []int{1, 1, 1}, target: 2, expect: 2},
		{input: []int{1, 2, 3}, target: 3, expect: 2},

	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCase {
			assert.Equal(t, test.expect, subarraySum(test.input, test.target))
		}
	})
}

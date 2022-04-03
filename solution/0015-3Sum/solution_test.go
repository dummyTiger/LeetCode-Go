package _015_3Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		input  []int
		expect [][]int
	}{
		{input: []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCases {
			assert.Equal(t, test.expect, threeSum(test.input))
		}
	})
}

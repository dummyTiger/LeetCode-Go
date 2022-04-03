package _209_MinimumSizeSubarraySum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		expect int
	}{
		{target: 7, nums: []int{2, 3, 1, 2, 4, 3},expect: 2},
	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCases {
			assert.Equal(t, test.expect, minSubArrayLen(test.target, test.nums))
		}
	})
}

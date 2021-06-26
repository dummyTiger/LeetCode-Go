package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 12
	res := TwoSum(nums, target)
	assert.Equal(t, 4, res[0])
	assert.Equal(t, 6, res[1])
}

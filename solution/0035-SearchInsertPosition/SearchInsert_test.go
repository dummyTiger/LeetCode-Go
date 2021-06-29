package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	nums1 := []int{1}
	assert.Equal(t, 2, searchInsert(nums, 5))
	assert.Equal(t, 1, searchInsert(nums, 2))
	assert.Equal(t, 4, searchInsert(nums, 7))
	assert.Equal(t, 0, searchInsert(nums, 0))
	assert.Equal(t, 0, searchInsert(nums1, 0))

}

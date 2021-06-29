package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums2 := []int{1, 1, 2}
	nums3 := []int{1, 1, 1}
	assert.Equal(t, 5, removeDuplicates(nums1))
	assert.Equal(t, 2, removeDuplicates(nums2))
	assert.Equal(t, 2, removeDuplicates(nums3))

}

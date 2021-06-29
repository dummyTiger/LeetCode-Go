package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	assert.Equal(t, 2, removeElement(nums, 3))
	assert.Equal(t, 5, removeElement(nums1, 2))

}

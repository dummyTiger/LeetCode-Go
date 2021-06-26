package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{0,0}
	num2 := []int{0,0}
	res := findMedianSortedArrays(num1, num2)

	assert.Equal(t, 0.0, res)
}

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {

	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 0, maxSubArray([]int{-1,0, -2}))

}

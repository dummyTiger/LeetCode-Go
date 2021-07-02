package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 16, maxArea([]int{4, 3, 2, 1, 4}))
	assert.Equal(t, 2, maxArea([]int{1,2,1}))
	assert.Equal(t, 1, maxArea([]int{1,1}))


}

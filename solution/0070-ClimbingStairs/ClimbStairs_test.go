package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t,2, climbStairs(2))
	assert.Equal(t,3, climbStairs(3))

}

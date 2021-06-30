package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(5))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 0, mySqrt(0))
	assert.Equal(t, 10, mySqrt(101))
}

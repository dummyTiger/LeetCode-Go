package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, true,isPowerOfTwo(8))
	assert.Equal(t, false,isPowerOfTwo(9))

}

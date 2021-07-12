package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	assert.Equal(t, true, isPowerOfThree(27))
	assert.Equal(t, false, isPowerOfThree(19))
}

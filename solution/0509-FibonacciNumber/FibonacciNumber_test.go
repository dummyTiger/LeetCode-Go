package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 1, fib(2))
	assert.Equal(t, 2, fib(3))

}

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestValidParentheses(t *testing.T)  {

	assert.Equal(t,2,longestValidParentheses("(()(()()"))
}

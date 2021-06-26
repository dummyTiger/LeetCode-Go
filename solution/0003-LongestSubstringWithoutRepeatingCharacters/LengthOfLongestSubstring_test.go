package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	assert.Equal(t,1,lengthOfLongestSubstring(s))
}

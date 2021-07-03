package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//动态规划
func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "aba", longestPalindrome("abaca"))
	assert.Equal(t, "aa", longestPalindrome("aab"))
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "aca", longestPalindrome("aackacaa"))

}

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("A man, a plan, a canal: Panama"))
	assert.Equal(t, false, isPalindrome("race a car"))
	assert.Equal(t, false, isPalindrome("OP"))


}

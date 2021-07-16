package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	//assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 1, lengthOfLastWord("a "))

}

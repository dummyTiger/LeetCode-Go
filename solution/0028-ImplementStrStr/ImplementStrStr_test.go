package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {

	haystack:="mississippi"

	needle := "issip"

	assert.Equal(t,strStr(haystack,needle),4)
}

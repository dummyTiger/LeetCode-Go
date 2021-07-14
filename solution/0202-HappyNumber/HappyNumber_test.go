package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsHappy(t *testing.T)  {
	assert.Equal(t, true,isHappy(19))
	assert.Equal(t, true,isHappy(91))
	assert.Equal(t, false,isHappy(2))
	assert.Equal(t, true,isHappy(7))



}

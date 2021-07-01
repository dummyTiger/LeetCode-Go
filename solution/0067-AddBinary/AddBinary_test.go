package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	//assert.Equal(t, "10", addBinary("1", "1"))
	//assert.Equal(t, "10101", addBinary("1010", "1011"))
	//assert.Equal(t, "110110", addBinary("100", "110010"))
	assert.Equal(t, "110001", addBinary("101111", "10"))


}

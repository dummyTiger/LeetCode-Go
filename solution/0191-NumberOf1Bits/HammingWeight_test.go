package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, hammingWeight(00000000000000000000000000001011), 3)
	assert.Equal(t, hammingWeight(00000000000000000000000010000000), 1)
	//assert.Equal(t, hammingWeight(11111111111111111111111111111101), 31)

}

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	assert.Equal(t, 249999,countPrimes1(500000))
}

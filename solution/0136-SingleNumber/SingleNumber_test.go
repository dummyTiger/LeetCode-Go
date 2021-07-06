package leecode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{4, 2, 1, 2, 1}))

}

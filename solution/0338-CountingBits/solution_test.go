package _338_CountingBits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBits(t *testing.T) {
	testCase := []struct {
		input  int
		expect []int
	}{
		{input: 5, expect: []int{0, 1, 1, 2, 1, 2}},
	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCase {
			assert.Equal(t, test.expect, countBits(test.input))
		}
	})
}

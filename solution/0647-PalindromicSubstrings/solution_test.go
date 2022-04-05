package _647_PalindromicSubstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		expect int
	}{
		// TODO: test cases
		{name: "abc", expect: 3},
		{name: "aaa", expect: 6},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, countSubstrings(test.name))
		})
	}
}

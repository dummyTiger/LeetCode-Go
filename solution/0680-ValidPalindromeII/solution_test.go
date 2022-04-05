package _680_ValidPalindromeII

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
	}{
		{name: "aba", expect: true},
		{name: "abc", expect: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, validPalindrome(test.name))
		})
	}
}

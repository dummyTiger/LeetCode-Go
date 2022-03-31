package _014_CheckInclusion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		s1     string
		s2     string
		expect bool
	}{
		{s1: "ab", s2: "eidbaooo", expect: true},
		{s1: "ab", s2: "eidboaoo", expect: false},
		{s1: "hello", s2: "ooolleoooleh", expect: false},
	}

	t.Run("t", func(t *testing.T) {
		for _, test := range testCases {
			assert.Equal(t, test.expect, checkInclusion(test.s1, test.s2))
		}
	})
}

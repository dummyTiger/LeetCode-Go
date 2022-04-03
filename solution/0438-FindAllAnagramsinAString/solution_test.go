package _438_FindAllAnagramsinAString

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		s1     string
		s2     string
		expect []int
	}{
		{s1: "eidbaooo", s2: "ab", expect: []int{3}},
		{s1: "eidboaoo", s2: "ab", expect: []int{}},
		{s1: "ooolleoooleh", s2: "hello", expect: []int{}},
	}

	t.Run("t", func(t *testing.T) {
		for _, test := range testCases {
			assert.Equal(t, test.expect, find(test.s1, test.s2))
		}
	})
}

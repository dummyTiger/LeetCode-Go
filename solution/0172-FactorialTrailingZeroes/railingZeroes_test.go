package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRailingZeroes(t *testing.T) {

	testCases := []struct {
		Input  int
		Expect int
	}{
		{Input: 10, Expect: 2},
		{Input: 11, Expect: 2},
		{Input: 13, Expect: 2},
		{Input: 14, Expect: 2},
		{Input: 15, Expect: 3},
		//{Input: 30, Expect: 7},
		{Input: 625, Expect: 156},
	}

	t.Run("test", func(t *testing.T) {

		for _, testCase := range testCases {

			assert.Equal(t, testCase.Expect, railingZeroes(testCase.Input))
		}

	})

}

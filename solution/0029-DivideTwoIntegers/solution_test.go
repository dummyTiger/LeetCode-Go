package _029_DivideTwoIntegers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	testCase := []struct {
		Dividend int
		Divisor  int
		Expect   int
	}{
		//{Dividend: 10, Divisor: 3, Expect: 3},
		//{Dividend: -7, Divisor: 3, Expect: -2},
		//{Dividend: 1, Divisor: 1, Expect: 1},
		{Dividend: 2147483647, Divisor: 1, Expect: 2147483647},
	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCase {
			assert.Equal(t, test.Expect, Solution(test.Dividend, test.Divisor))
		}
	})
}

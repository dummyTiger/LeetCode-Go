package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	//assert.Equal(t, "10", addBinary("1", "1"))
	//assert.Equal(t, "10101", addBinary("1010", "1011"))
	//assert.Equal(t, "110110", addBinary("100", "110010"))
	assert.Equal(t, "110001", addBinary("101111", "10"))

}

func TestAddBinary1(t *testing.T) {
	testCase := []struct {
		input1 string
		input2 string
		expect string
	}{
		{input1: "101111", input2: "10", expect: "110001"},
	}

	t.Run("test", func(t *testing.T) {
		for _, test := range testCase {
			assert.Equal(t, test.expect, addBinary1(test.input1, test.input2))
		}
	})
}

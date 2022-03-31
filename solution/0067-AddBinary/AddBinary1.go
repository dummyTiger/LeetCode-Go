package leetcode

import (
	"strconv"
	"strings"
)

func addBinary1(a string, b string) string {

	leni, lenj := len(a), len(b)

	res := make([]string, 0)

	carry := uint8(0)

	for i, j := leni-1, lenj-1; i >= 0 || j >= 0; {
		ii, jj := uint8(0), uint8(0)
		if i >= 0 {
			ii = a[i] - '0'
		} else {
			ii = 0
		}

		if j >= 0 {
			jj = b[j] - '0'
		} else {
			jj = 0
		}

		sum := ii + jj + carry

		if sum >= 2 {
			carry = 1
			res = append(res, strconv.Itoa(int(sum-2)))
		} else {
			carry = 0
			res = append(res, strconv.Itoa(int(sum)))
		}
		i--
		j--
	}

	if carry == 1 {
		res = append(res, "1")
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, "")
}

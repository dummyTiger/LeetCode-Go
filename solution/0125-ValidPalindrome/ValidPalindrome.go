package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	flag := true
	s = strings.ToUpper(s)
	for i, j := 0, len(s)-1; i < j; {
		if !((s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90)) {
			i++
			continue
		}
		if !((s[j] >= 48 && s[j] <= 57) || (s[j] >= 65 && s[j] <= 90)) {
			j--
			continue
		}

		if s[i] == s[j] {
			i++
			j--
			continue
		} else {
			flag = false
			break
		}
	}

	return flag
}

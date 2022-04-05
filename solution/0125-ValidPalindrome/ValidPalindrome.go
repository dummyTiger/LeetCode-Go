package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	flag := true
	s = strings.ToUpper(s)
	for i, j := 0, len(s)-1; i < j; {
		if !IsLetterOrNumber(rune(s[i])) {
			i++
			continue
		}
		if !IsLetterOrNumber(rune(s[j])) {
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
func IsLetterOrNumber(v rune) bool {
	return unicode.IsNumber(v) || unicode.IsLetter(v)
}

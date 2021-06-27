package leetcode

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	begin, end := 0, len(str)-1
	for begin < end {
		if str[begin] != str[end] {
			return false
		} else {
			begin += 1
			end -= 1
		}
	}
	return true
}

package _680_ValidPalindromeII

func validPalindrome(s string) bool {
	begin := 0
	end := len(s) - 1

	for ; begin < len(s)/2; {
		if s[begin] != s[end] {
			break
		}
		begin++
		end--
	}

	return begin == len(s)/2 || isValidPalindrome(s, begin, end-1) ||
		isValidPalindrome(s, begin+1, end)

}

func isValidPalindrome(s string, begin, end int) bool {
	for ; begin < end; {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return begin >= end
}

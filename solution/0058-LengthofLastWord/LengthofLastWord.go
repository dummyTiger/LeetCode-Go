package leetcode

func lengthOfLastWord(s string) int {
	size := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			size++
		} else {
			if size != 0 {
				break
			}
		}
	}
	return size
}

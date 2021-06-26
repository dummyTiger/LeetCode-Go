package leetcode

func lengthOfLongestSubstring(s string) int {
	res, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		res = 0
		tempMap := make(map[uint8]bool)
		for j := i; j < len(s); j++ {
			b := s[j]
			if _, ok := tempMap[b]; !ok {
				tempMap[b] = true
				res = res + 1
				if res > maxLen {
					maxLen = res
				}
			} else {
				break
			}
		}
	}
	return maxLen
}

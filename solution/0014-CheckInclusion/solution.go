package _014_CheckInclusion

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	counts := [26]int{}
	for i := 0; i < len(s1); i++ {
		counts[s1[i]-'a']++
		counts[s2[i]-'a']--
	}
	if isAllZero(counts) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		counts[s2[i]-'a']--
		counts[s2[i-len(s1)]-'a']++
		if isAllZero(counts) {
			return true
		}
	}
	return false

}

func isAllZero(arr [26]int) bool {
	for _, i := range arr {
		if i != 0 {
			return false
		}
	}
	return true
}

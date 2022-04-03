package _438_FindAllAnagramsinAString

func find(s1 string, s2 string) []int {
	if len(s1) < len(s2) {
		return []int{}
	}

	res := make([]int, 0)
	counts := [26]int{}
	i := 0
	for ; i < len(s2); i++ {
		counts[s2[i]-'a']++
		counts[s1[i]-'a']--
	}
	if isAllZero(counts) {
		res = append(res, 0)
	}
	for ; i < len(s1); i++ {
		counts[s1[i]-'a']--
		counts[s1[i-len(s2)]-'a']++
		if isAllZero(counts) {
			res = append(res, i-len(s2)+1)
		}
	}
	return res

}

func isAllZero(arr [26]int) bool {
	for _, i := range arr {
		if i != 0 {
			return false
		}
	}
	return true
}

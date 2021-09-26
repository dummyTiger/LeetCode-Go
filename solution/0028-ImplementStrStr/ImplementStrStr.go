package leetcode

func strStr(haystack string, needle string) int {

	if haystack=="" && needle!="" {
		return -1
	}

	if needle == "" {
		return 0
	}

	if len(haystack)<len(needle) {
		return  -1
	}

	longIndex:=0
	shortIndex:=0
	currentIndex:=-1

	for ;longIndex<len(haystack)&&shortIndex<len(needle); {
		if haystack[longIndex] == needle[shortIndex] {
			if shortIndex == 0 {
				currentIndex = longIndex
			}
			longIndex  = longIndex +1
			shortIndex = shortIndex +1
		} else {
			if currentIndex!=-1 {
				longIndex = currentIndex+1
				currentIndex = -1
				shortIndex = 0
			}else {
				longIndex = longIndex +1
				shortIndex = 0
				currentIndex = -1
			}
		}
	}

	if shortIndex!=len(needle) {
		currentIndex = -1
	}

	return currentIndex
}

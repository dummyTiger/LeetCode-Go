package _647_PalindromicSubstrings

//前面都是从字符串的两端开始向里移动指针来判断字符串是否是一个回文，
//其实也可以换一个方向从字符串的中心开始向两端延伸。如果存在一个长度为m的回文子字符串，
//则分别再向该回文的两端延伸一个字符，并判断回文前后的字符是否相同。
//如果相同，就找到了一个长度为m+2的回文子字符串。
//例如，在字符串"abcba"中，从中间的"c"出发向两端延伸一个字符，由于"c"前后都是字符'b'，
//因此找到了一个长度为3的回文子字符串"bcb"。然后向两端延伸一个字符，由于"bcb"的前后都是字符'a'，
//因此又找到一个长度为5的回文子字符串"abcba"。
//值得注意的是，回文的长度既可以是奇数，也可以是偶数。长度为奇数的回文的对称中心只有一个字符，
//而长度为偶数的回文的对称中心有两个字符
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	c := 0

	for i := 0; i < len(s); i++ {
		c += count(s, i, i)
		c += count(s, i, i+1)
	}
	return c

}

func count(s string, start, end int) int {
	c := 0
	for ; start >=0 && end < len(s) && s[start] == s[end]; {
		c++
		start--
		end++
	}
	return c
}

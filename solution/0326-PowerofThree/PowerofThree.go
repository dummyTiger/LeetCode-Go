package leetcode

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	flag := true
	for {
		if n == 1 {
			break
		}
		if n%3 == 0 {
			n = n / 3
		} else {
			flag = false
			break
		}
	}
	return flag
}

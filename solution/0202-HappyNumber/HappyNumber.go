package leetcode

//todo 暴力解法
func isHappy(n int) bool {
	res := 0
	time := 500
	for ; time > 0; time = time - 1 {
		if n == 1 {
			res = n
			break
		}
		res = 0
		for {
			if n == 0 {
				n = res
				break
			}
			left := n % 10
			res += left * left
			n = n / 10
		}
	}

	return res == 1
}

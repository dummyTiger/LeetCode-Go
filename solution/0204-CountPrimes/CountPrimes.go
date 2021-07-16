package leetcode

func countPrimes1(n int) int {
	count := 0
	res := make([]bool, n)

	for i := 2; i < n; i++ {
		if !res[i] {
			count++
			for j := i + 1; j < n; j = j + i {
				res[j] = true
			}
		}
	}
	return  count
}

//超时
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	res := make([]int, 0)
	begin := []int{0, 0, 1, 2, 2, 3, 3, 4}
	res = append(res, begin...)
	pri := make([]int, 0)
	a := []int{2, 3, 5, 7}
	pri = append(pri, a...)
	for i := 8; i <= n; i++ {
		flag := false
		for _, p := range pri {
			if i%p == 0 {
				flag = true
				break
			}
		}

		if flag {
			res = append(res, res[i-1])
		} else {
			res = append(res, res[i-1]+1)
			pri = append(pri, i)
		}
	}
	return res[n-1]
}

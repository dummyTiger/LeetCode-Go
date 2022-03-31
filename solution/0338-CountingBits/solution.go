package _338_CountingBits

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		j := i

		for {
			if j != 0 {
				res[i] = res[i] + 1
				j = j & (j - 1)
			} else {
				break
			}
		}
	}

	return res
}

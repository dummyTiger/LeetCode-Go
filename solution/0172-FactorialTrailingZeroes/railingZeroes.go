package leetcode

func railingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	factor := 1
	numOfZero := 0

	for i := 1; i <= n; i++ {
		factor = factor * i
		for {
			if factor%10 == 0 {
				numOfZero++
				factor = factor / 10
			} else {
				factor = factor % 10
				break
			}
		}
	}
	return numOfZero
}

package leetcode

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	if x == 0 {
		return 0
	}

	res := 1

	for i := 1; i <= x/2; i++ {
		temp := i * i
		if temp < x {
			res = i
			continue
		} else if temp == x {
			res = i
			break
		} else {
			res = i - 1
			break
		}
	}
	return res
}

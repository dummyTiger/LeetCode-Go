package leetcode

import "math"

func reverse(x int) int {
	temp := make([]int, 0)
	for x != 0 {
		temp = append(temp, x%10)
		x = x / 10
	}
	var result int64
	time := 1
	for i := len(temp) - 1; i >= 0; i-- {
		result = result + int64(temp[i]*time)
		time = time * 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return int(result)

}

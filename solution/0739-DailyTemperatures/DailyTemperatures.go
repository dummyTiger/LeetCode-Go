package leetcode

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, 0)

	if len(temperatures) == 1 {
		return []int{0}
	}

	for i := 0; i < len(temperatures); i++ {
		flag := false
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res = append(res, j-i)
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, 0)
		}
	}
	return res
}

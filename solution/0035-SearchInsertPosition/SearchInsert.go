package leetcode

func searchInsert(nums []int, target int) int {
	position := 0
	flag := false
	for idx, value := range nums {
		if value < target {
			continue
		} else {
			position = idx
			flag = true
			break
		}
	}

	if !flag {
		position = len(nums)
	}

	return position
}

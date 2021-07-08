package leetcode

func moveZeroes(nums []int) {
	skip := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			skip++
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == 0 {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	idx := 0
	for i, v := range nums {
		if v == 0 {
			idx = i - 1
			break
		}
	}

	for i, j := skip, idx; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func moveZeroes1(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			continue
		}

		nums[slow] = nums[fast]
		slow++
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

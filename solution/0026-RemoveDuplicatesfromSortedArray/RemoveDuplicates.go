package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	length := 1
	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		if nums[i]^nums[j] == 0 {
			j++
		} else {
			nums[length] = nums[j]
			i = j
			length++
		}
	}
	nums = nums[0:length]
	return length
}

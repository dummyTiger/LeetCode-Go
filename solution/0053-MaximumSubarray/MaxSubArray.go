package leetcode

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]

	res := 0

	for i := 0; i <= len(nums)-1; i++ {
		res = 0
		for j := i; j <= len(nums)-1; j++ {
			res = res + nums[j]
			if res > max {
				max = res
			}
		}

	}
	return max
}

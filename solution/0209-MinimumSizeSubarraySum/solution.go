package _209_MinimumSizeSubarraySum

import "math"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLength := math.MaxInt32
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for ; left <= right && sum >= target; {
			if minLength > right-left+1 {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	} else {
		return minLength
	}

}

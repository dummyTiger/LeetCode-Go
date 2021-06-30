package leetcode

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	len:=len(nums)
	return nums[len-1] * nums[len-2] -  nums[0]*nums[1]
}


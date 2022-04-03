package _015_3Sum

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for p, i := range nums {
		n := make([]int, 0)
		n = append(n, nums[:p]...)
		n = append(n, nums[p+1:]...)
		result := helper(n, 0-i)
		if len(result) != 0 {
			result = append(result, i)
			res = append(res, result)
		}
	}

	return res
}

func helper(input []int, target int) []int {
	for i, j := 0, len(input)-1; i < j; {
		sum := input[i] + input[j]
		if sum == target {
			return []int{input[i], input[j]}
		} else if sum < target {
			i = i + 1
		} else {
			j = j - 1
		}
	}
	return []int{}
}

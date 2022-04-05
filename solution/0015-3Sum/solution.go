package _015_3Sum

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	i := 0
	for i < len(nums)-2 {
		helper(nums, i, &res)
		temp := nums[i]
		for i < len(nums) && nums[i] == temp {
			i++
		}
	}

	return res
}

func helper(input []int, i int, res *[][]int) {
	j := i + 1
	k := len(input) - 1
	for j < k {
		sum := input[i] + input[j] + input[k]
		if sum == 0 {
			*res = append(*res, []int{input[i], input[j], input[k]})
			temp := input[j]
			for j < k && input[j] == temp {
				j++
			}
		} else if sum < 0 {
			j++
		} else {
			k--
		}
	}
}

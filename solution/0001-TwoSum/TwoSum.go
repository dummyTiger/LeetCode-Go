package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, k}
		}
		m[v] = k
	}
	return nil
}

package _167_TwoSumII

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i = i + 1
		} else {
			j = j - 1
		}
	}
	return []int{}
}

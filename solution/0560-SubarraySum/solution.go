package _560_SubarraySum

func subarraySum(nums []int, k int) int {
	length := len(nums)
	sum := 0
	count := 0
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
		sum = 0
	}

	return count
}

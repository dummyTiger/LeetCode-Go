package _713_NumSubarrayProductLessThanK

func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	product := 1
	count := 0
	for right := 0; right < len(nums); right++ {
		product = product * nums[right]
		for ; left <= right && product >= k; {
			product /= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count
}

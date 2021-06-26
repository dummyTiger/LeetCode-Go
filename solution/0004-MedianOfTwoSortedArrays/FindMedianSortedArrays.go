package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	len := len(nums1)
	if len%2 == 0 {
		return float64(nums1[len/2-1]+nums1[len/2]) / 2.0
	} else {
		return float64(nums1[len/2])
	}
}

package leetcode

func maxArea(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1
	for i < j {

		w := getMinValue(height[i], height[j])
		area := w * (j - i)

		if area > maxArea {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return maxArea
}

func getMinValue(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

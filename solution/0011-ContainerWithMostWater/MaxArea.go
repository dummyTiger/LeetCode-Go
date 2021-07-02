package leetcode

//todo unsolved
func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i <= len(height)-2; i++ {
		if height[i] == 0 {
			continue
		}
		for j := i + 1; j <= len(height)-1; j++ {
			if height[j] == 0 {
				continue
			}
			w := j - i
			h := 0
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			area := w * h
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

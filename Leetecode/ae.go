package Leetecode

func maxArea(height []int) int {
	var left, right = 0, len(height) - 1
	var result int
	for left < right {
		curArea := (right - left) * min(height[left], height[right])
		result = max(result, curArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return result
}

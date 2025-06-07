package Leetecode

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	var ans int
	for left < right && left < len(height) && right > 0 {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

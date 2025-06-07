package Leetecode

func moveZeroes(nums []int) {
	var left, right, n = 0, 0, len(nums)

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

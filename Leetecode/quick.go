package Leetecode

func findKthLargest(nums []int, k int) int {
	return quickselect(nums, 0, len(nums)-1, k-1)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	i, j := l-1, r+1
	privot := nums[l]

	for i < j {
		for i++; nums[i] > privot; i++ {
		}
		for j--; nums[j] < privot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

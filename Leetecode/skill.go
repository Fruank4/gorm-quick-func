package Leetecode

func majorityElement(nums []int) int {
	current, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			current = nums[i]
			count = 1
		} else {
			if nums[i] == current {
				count++
			} else {
				count--
			}
		}

	}
	return current
}

func sortColors(nums []int) {
	// 2,0,2,1,1,0
	// 0,0,1,1,2,2

	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

func nextPermutation(nums []int) {
	// 1, 2, 4, 1
	// 1, 4, 2, 1
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// 1.找到合适的i和j
	for nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] > nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j = j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0

	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

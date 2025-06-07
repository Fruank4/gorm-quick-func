package Leetecode

import "sort"

/*
*
[-1,0,1,2,-1,-4]

-4， -1， -1 ，0， 1， 2
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		var left, right = k + 1, len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == -nums[k] {
				result = append(result, []int{nums[k], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] < -nums[k] {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

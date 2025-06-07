package Leetecode

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	// set
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	var longest = 0
	for num, _ := range numSet {
		if !numSet[num-1] {
			var currentNum = num
			var currentLen = 1
			for numSet[currentNum+1] {
				currentNum++
				currentLen++
			}
			if currentLen > longest {
				longest = currentLen
			}
		}
	}
	return longest
}

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	sort.Ints(nums)

	// 1, 3, 4, 5, 6, 7, 8, 9
	var maxLen = 0
	var curLen = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			curLen++
		} else if nums[i] == nums[i-1] {

		} else {
			curLen = 1
		}

		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

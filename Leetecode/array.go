package Leetecode

import (
	"math"
	"sort"
)

// [-2,1,-3,4,-1,2,1,-5,4]
func maxSubArray(nums []int) int {
	preSum := 0
	minPreSum := 0
	ans := math.MinInt
	for _, num := range nums {
		preSum += num
		ans = max(ans, preSum-minPreSum)
		if minPreSum > preSum {
			minPreSum = preSum
		}
	}

	return ans
}

/**
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	for _, interval := range intervals {
		if len(ans) > 0 {
			cur := ans[len(ans)-1]
			if cur[1] >= interval[0] {
				cur[1] = max(cur[1], interval[0])
			} else {
				ans = append(ans, interval)
			}
		} else {
			ans = append(ans, interval)
		}
	}
	return ans
}

/*
*
输入: nums = [1,2,3,4,5,6], k = 3
输出: 4 5 6 1 2 3
*/
func rotate(nums []int, k int) {
	n := len(nums)
	for start, count := 0, gcd(n, k); start < count; start++ {
		cur := start
		pre := nums[start]
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	if a != 0 {
		a, b = b%a, a
	}
	return b
}

/**
输入: nums = [1,2,3,4]
输出: [24,12,8,6]
*/

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	sufProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * sufProduct
		sufProduct = sufProduct * nums[i]
	}
	return ans
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < len(nums); i++ {
		value := abs(nums[i])
		if value <= n {
			nums[value-1] = -abs(nums[value-1])
		}
	}
	var i int
	for i = 0; nums[i] < 0; {
		i++
	}
	return i + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package Leetecode

// [7,1,5,3,6,4]
//

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var minValue = prices[0]
	var ans = 0
	for i := 1; i < len(prices); i++ {
		if minValue > prices[i] {
			minValue = prices[i]
		} else {
			ans = max(ans, prices[i]-minValue)
		}
	}
	return ans
}

func maxProfit2(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	dp := make([]int, len(prices))
	dp[0] = 0

	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], dp[i-1]+prices[i]-prices[i-1])
	}
	return dp[len(prices)-1]
}

// nums = [2,3,1,1,4]
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var limit = nums[0]
	for i := 1; i <= limit; i++ {
		limit = max(limit, i+nums[i])
		if limit >= len(nums)-1 {
			return true
		}
	}

	return false
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// 每次突破都是跳跃时刻
	ans := 0
	limit := 0
	maxPosition := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == limit {
			limit = maxPosition
			ans++
		}
	}
	return ans
}

func partitionLabels(s string) []int {
	mp := make(map[int32]int)
	for i, c := range s {
		mp[c-'a'] = i
	}

	end := 0
	start := 0
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		end = max(end, mp[int32(s[i]-'a')])
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

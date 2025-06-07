package Leetecode

func subarraySum(nums []int, k int) int {
	pre, ans := 0, 0
	mp := make(map[int]int)
	mp[0] = 1
	for _, num := range nums {
		pre += num
		ans += mp[pre-k]
		mp[pre]++
	}
	return ans
}

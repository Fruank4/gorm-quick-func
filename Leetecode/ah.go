package Leetecode

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]bool)
	ans := 0
	var right = -1
	for left := 0; left < len(s); left++ {
		if left != 0 {
			delete(mp, s[left-1])
		}

		for right+1 < len(s) && !mp[s[right+1]] {
			mp[s[right+1]] = true
			right++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

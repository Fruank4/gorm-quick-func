package Leetecode

import "maps"

func FindAnagrams(s string, p string) []int {
	// 初始化
	ans := make([]int, 0)
	if len(s) < len(p) {
		return ans
	}
	pMap := make(map[byte]int)
	sMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
		if i < len(p)-1 {
			sMap[s[i]]++
		}
	}

	for left, right := 0, len(p)-1; right < len(s); right++ {
		sMap[s[right]]++
		if maps.Equal(sMap, pMap) {
			ans = append(ans, left)
		}
		sMap[s[left]]--
		if sMap[s[left]] == 0 {
			delete(pMap, s[left])
		}
		left++
	}
	return ans
}

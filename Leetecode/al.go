package Leetecode

import "math"

// ADOBECODEBANC", t = "ABC" "BANC"
// ABC
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	check := func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	ansL, ansR := -1, -1
	lenth := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if tMap[s[r]] > 0 {
			sMap[s[r]]++
		}
		for check() && l <= r {
			if r-l < lenth {
				lenth = r - l
				ansL, ansR = l, r+1
			}

			if tMap[s[l]] > 0 {
				sMap[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}

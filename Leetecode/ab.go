package Leetecode

/*
*

	strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

	输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]int][]string)
	for _, str := range strs {
		counts := [26]int{}
		for _, c := range str {
			counts[c-'a']++
		}
		strMap[counts] = append(strMap[counts], str)
	}
	ans := make([][]string, 0)
	for _, value := range strMap {
		ans = append(ans, value)
	}
	return ans
}

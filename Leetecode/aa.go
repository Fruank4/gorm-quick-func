package Leetecode

func twoSum(nums []int, target int) []int {
	var num2Index = make(map[int]int)
	for i, num := range nums {
		value, exist := num2Index[target-num]
		if exist {
			return []int{i, value}
		} else {
			num2Index[num] = i
		}
	}

	return []int{0}
}

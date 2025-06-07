package Leetecode

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	push := func(num int) {
		for len(queue) > 0 && num >= queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, len(nums)-k+1)
	ans[0] = queue[0]
	for i := k; i < len(nums); i++ {
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		push(nums[i])
		ans = append(ans, queue[0])
	}
	return ans
}

package Leetecode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans = make([]int, 0)
	inorder(root)
	return ans
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	ans = append(ans, root.Val)
	inorder(root.Right)
	return
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(rootA *TreeNode, rootB *TreeNode) bool {
	if rootA == nil && rootB == nil {
		return true
	}
	if rootA == nil || rootB == nil {
		return false
	}

	return rootA.Val == rootB.Val && isMirror(rootA.Left, rootB.Right) && isMirror(rootA.Right, rootB.Left)
}

var result = 0

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	diameterCode(root)
	return result
}

func diameterCode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := diameterCode(root.Left)
	right := diameterCode(root.Right)
	result = max(result, left+right)
	return 1 + max(left, right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	return 1 + max(leftMaxDepth, rightMaxDepth)
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	offer := func(node *TreeNode) {
		if node != nil {
			queue = append(queue, node)
		}
	}
	poll := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}

	offer(root)
	ans := make([][]int, 0)
	for len(queue) > 0 {
		count := len(queue)
		subAns := make([]int, 0)
		for i := 0; i < count; i++ {
			node := poll()
			subAns = append(subAns, node.Val)
			offer(node.Left)
			offer(node.Right)
		}
		ans = append(ans, subAns)
	}
	return ans
}

/*
*
[-10,-3,0,5,9]
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	if mid+1 > len(nums)-1 {
		root.Right = nil
	} else {
		root.Right = sortedArrayToBST(nums[mid+1 : len(nums)-1])
	}
	return root
}

var list []*TreeNode

func isValidBSTCore(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftIs := isValidBSTCore(root.Left)
	if !leftIs {
		return false
	}
	if len(list) == 0 {
		list = append(list, root)
	} else {
		before := list[len(list)-1]
		if before.Val >= root.Val {
			return false
		} else {
			list = append(list, root)
		}
	}

	rightIs := isValidBSTCore(root.Right)
	if !rightIs {
		return false
	}
	return true
}

var maxPathSumAns = math.MinInt32

func maxPathSum(root *TreeNode) int {
	maxPathSumAns = 0
	maxPathSumRecur(root)
	return maxPathSumAns
}

func maxPathSumRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxPathSumRecur(root.Left)
	rightMax := maxPathSumRecur(root.Right)

	maxPathSumAns = max(maxPathSumAns, leftMax+rightMax+root.Val)
	return max(0, root.Val+max(leftMax, rightMax))
}

func rightSideView(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	offer := func(node *TreeNode) {
		if node != nil {
			queue = append(queue, node)
		}
	}
	poll := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}

	offer(root)
	result := make([]int, 0)
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			node := poll()
			if i == count-1 {
				result = append(result, node.Val)
			}
			offer(node.Left)
			offer(node.Right)
		}
	}
	return result
}

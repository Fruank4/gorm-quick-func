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
	maxPathSumAns = math.MinInt32
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if (leftNode == p && rightNode == q) || (rightNode == p && leftNode == q) {
		return root
	}
	if leftNode != nil {
		return leftNode
	}
	if rightNode != nil {
		return rightNode
	}

	return nil
}

var pre *TreeNode

func flatten(root *TreeNode) {
	pre = nil
	flattenCore(root)
}

func flattenCore(root *TreeNode) {
	pre = nil
	if root == nil {
		return
	}
	flattenCore(root.Right)
	flattenCore(root.Left)
	root.Left = nil
	root.Right = pre
	pre = root
}

var TargetSum int64
var pathCountMap map[int64]int

func pathSum(root *TreeNode, targetSum int) int {
	TargetSum = int64(targetSum)
	pathCountMap = make(map[int64]int)
	result = 0
	pathCountMap[0]++
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, val int64) {
	if root == nil {
		return
	}
	curVal := int64(root.Val) + val
	needVal := curVal - TargetSum

	result += pathCountMap[needVal]

	pathCountMap[curVal]++
	dfs(root.Left, curVal)
	dfs(root.Right, curVal)
	pathCountMap[curVal]--
}

/**
preorder [3, 9,10, 20,15,7], inorder = [9,10, 3,15,20,7]
*/

var Pre []int
var Ino []int

func buildTree(preorder []int, inorder []int) *TreeNode {
	Pre = preorder
	Ino = inorder
	return buildTreeCore(0, len(Pre)-1, 0, len(Ino)-1)
}

func buildTreeCore(preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preEnd < preStart || inEnd < inStart {
		return nil
	}
	root := &TreeNode{
		Val: Pre[preStart],
	}
	var i = inStart
	for i < inEnd && Ino[i] != Pre[preStart] {
		i++ //2
	}

	root.Left = buildTreeCore(preStart+1, preStart+i-inStart, inStart, i-1)
	root.Right = buildTreeCore(preStart+i-inStart+1, preEnd, i+1, inEnd)
	return root
}

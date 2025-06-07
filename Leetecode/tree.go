package Leetecode

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

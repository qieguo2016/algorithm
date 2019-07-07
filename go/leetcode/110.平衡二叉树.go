/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
  if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	max := maxLeft
	if maxRight > maxLeft {
		max = maxRight
	}
	return 1 + max
}

func isBalanced(root *TreeNode) bool {
  if root == nil {
		return true
	}
	diff := maxDepth(root.Left) - maxDepth(root.Right)
	if diff > 1 || diff < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}


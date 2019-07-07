/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
	if root.Left == nil && root.Right == nil {
		return 1
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	max := maxLeft
	if maxRight > maxLeft {
		max = maxRight
	}
	return 1 + max
}


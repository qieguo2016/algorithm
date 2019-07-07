/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func isSymmetricByNode(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isSymmetricByNode(left.Left, right.Right) && isSymmetricByNode(left.Right, right.Left)
 }

func isSymmetric(root *TreeNode) bool {
  if root == nil {
		return true
	}
	return isSymmetricByNode(root.Left, root.Right)
}


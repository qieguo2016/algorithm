/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (35.24%)
 * Likes:    133
 * Dislikes: 0
 * Total Accepted:    7.4K
 * Total Submissions: 21.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空二叉树，返回其最大路径和。
 * 
 * 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
 * 
 * 示例 1:
 * 
 * 输入: [1,2,3]
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 * 
 * 输出: 6
 * 
 * 
 * 示例 2:
 * 
 * 输入: [-10,9,20,null,null,15,7]
 * 
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 * 
 * 输出: 42
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(root *TreeNode, res *int) int {
	if (root == nil) {
		return 0
	}
	left := max(helper(root.Left, res), 0)
	right := max(helper(root.Right, res), 0);
	*res = max(*res, left + right + root.Val);
	return max(left, right) + root.Val;
}
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}


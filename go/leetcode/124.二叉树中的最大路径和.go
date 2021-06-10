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

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 如果没有子函数的话，顶层就会直接返回左或右连接根节点
	// 增加子函数之后，就会比较左右单连根节点、左右都连根节点
	s := &solution{res: root.Val}
	s.helper(root)
	return s.res
}

// 全局变量会被上次结果影响，所以需要写到结构体内
type solution struct {
	res int
}

func (s *solution) helper(root *TreeNode) int {
	// 取左右子树的时候，只有两个情况：
	// 1. 不连接，直接对比各个子树的大小，所以用一个全局res来比较
	// 2. 连接，这种情况下只能连接终点为子树根节点的路径
	// 又因为要先算出左右子树，所以是后序遍历
	l, r := 0, 0
	if root.Left != nil {
		l = max(s.helper(root.Left), 0) // 小于0的子树直接丢弃
	}
	if root.Right != nil {
		r = max(s.helper(root.Right), 0) // 小于0的子树直接丢弃
	}
	s.res = max(s.res, l+r+root.Val)
	return max(l, r) + root.Val // 终点为子树根节点
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


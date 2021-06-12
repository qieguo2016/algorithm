/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (64.85%)
 * Likes:    478
 * Dislikes: 0
 * Total Accepted:    112.5K
 * Total Submissions: 173.4K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 * 示例:
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 层次遍历，每层取最后一个节点即可。
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []*TreeNode{}
		var last int
		for _, node := range queue {
			last = node.Val
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		ret = append(ret, last)
		queue = level
	}
	return ret
}

// @lc code=end


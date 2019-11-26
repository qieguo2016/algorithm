/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 *
 * https://leetcode-cn.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (62.64%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 16.2K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给出一个完全二叉树，求出该树的节点个数。
 *
 * 说明：
 *
 *
 * 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第
 * h 层，则该层包含 1~ 2^h 个节点。
 *
 * 示例:
 *
 * 输入:
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 *
 * 输出: 6
 *
 */

// package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 对比左右节点，找到满的一侧（个数为2^n-1），然后递归求解不满的一侧，加上顶点就是总数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getLevel(root.Left)
	r := getLevel(root.Right)
	if l == r { // 左侧必然是满的，右侧才会有部分节点在最低层
		lc := 0
		if l > 0 {
			lc = 2<<(l-1) - 1
		}
		return 1 + lc + countNodes(root.Right) // 顶点+左侧+递归右侧
	}
	// 右侧必然是满的
	rc := 0
	if r > 0 {
		rc = 2<<(r-1) - 1
	}
	return 1 + rc + countNodes(root.Left)
}

func getLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getLevel(root.Left)
	r := getLevel(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (67.33%)
 * Likes:    161
 * Dislikes: 0
 * Total Accepted:    31.1K
 * Total Submissions: 46K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]  
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3 
 * 
 * 输出: [3,2,1]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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

func travel(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	travel(root.Left, res)
	travel(root.Right, res)
	*res = append(*res, root.Val)
}
func postorderTraversal(root *TreeNode) []int {
  res := []int{}
	travel(root, &res)
	return res
}


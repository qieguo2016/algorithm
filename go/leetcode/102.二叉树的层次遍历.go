/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (56.18%)
 * Likes:    207
 * Dislikes: 0
 * Total Accepted:    27K
 * Total Submissions: 48.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层次遍历结果：
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
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

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	q := []*TreeNode{}
	if root != nil {
		q = append(q, root)
	}
	i := 0  // dequeue index
	j := len(q) // enqueue index
	k := 0
	for i < j {
		// dequeue
		node := q[i]
		if k >= len(ret) {
			ret = append(ret, []int{})
		}
		ret[k] = append(ret[k], node.Val)
		i++
		// enqueue
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if i >= j {
			k++  // level down
			j = len(q) // update enqueue index
		}
	}
	return ret
}


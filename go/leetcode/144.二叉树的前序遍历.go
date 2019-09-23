/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (61.61%)
 * Likes:    148
 * Dislikes: 0
 * Total Accepted:    41K
 * Total Submissions: 66.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 * 
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

// 递归
// func travel(root *TreeNode, res *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	*res = append(*res, root.Val)
// 	travel(root.Left, res)
// 	travel(root.Right, res)
// }

// func preorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	travel(root, &res)
// 	return res
// }

// 非递归
type Stack struct {
	Value []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.Value = append(s.Value, node)
}

func (s *Stack) Pop() (*TreeNode, bool) {
	if len(s.Value) == 0 {
		return nil, false
	}
	ret := s.Value[len(s.Value) - 1]
	s.Value = s.Value[:len(s.Value) -1]
	return ret, true
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	s := &Stack{}
	s.Push(root)
	for {
		node, ok := s.Pop()
		if !ok {
			break
		}
		res = append(res, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return res
}


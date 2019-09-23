/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (67.84%)
 * Likes:    280
 * Dislikes: 0
 * Total Accepted:    61.2K
 * Total Submissions: 89.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
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
 * 输出: [1,3,2]
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

// // 递归
// func travel(root *TreeNode, res *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	travel(root.Left, res)
// 	*res = append(*res, root.Val)
// 	travel(root.Right, res)
// }

// func inorderTraversal(root *TreeNode) []int {
// 	ret := []int{}
// 	travel(root, &ret)
// 	return ret
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

func (s *Stack) IsEmpty() bool {
	return len(s.Value) == 0
}

func inorderTraversal(root *TreeNode) []int {	
	res := []int{}
	if root == nil {
		return res
	}
	s := &Stack{}
	for !s.IsEmpty() || root != nil {
		if root != nil {
			s.Push(root)
			root = root.Left
		} else {
			node, ok := s.Pop()
			if !ok {
				continue
			}
			res = append(res, node.Val)
			root = node.Right
		}
	}
	return res
}

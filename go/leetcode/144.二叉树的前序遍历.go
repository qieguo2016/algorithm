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

func (s *Stack) IsEmpty() bool {
	return len(s.Value) == 0
}

// 非递归与层序遍历有点类似，第一次遇到节点就加入结果集，而且是从左到右的顺序
// 那么在同时拥有左右节点的时候，左节点是可以直接输出的，但是右节点需要存储起来，而且是底层节点先出，也就是后进先出
// 所以，需要用一个栈来存储右节点
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	s := &Stack{}
	for !s.IsEmpty() || root != nil {
		if root != nil {
			res = append(res, root.Val)
			if root.Right != nil {
				s.Push(root.Right)
			}
			root = root.Left
		} else {
			root, _ = s.Pop()
		}
	}
	return res
}


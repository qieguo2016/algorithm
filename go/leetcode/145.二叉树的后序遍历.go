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

// func travel(root *TreeNode, res *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	travel(root.Left, res)
// 	travel(root.Right, res)
// 	*res = append(*res, root.Val)
// }

// func postorderTraversal(root *TreeNode) []int {
//   res := []int{}
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

// 后序遍历是第三次遇到节点才输出，整体的规律是左右根的输出方式，解法的核心点是从右侧倒推
// 右节点压栈但是不输出，同时如果有左节点也压到另外一个栈，然后不断往右节点下层走，
// 到右节点尽头之后再从左节点栈取出来继续按照右节点深层递归的方式处理
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	// 从右边开始压栈
	right := &Stack{}
	left := &Stack{}
	for !left.IsEmpty() || root != nil {
		if root != nil {
			right.Push(root)
			if root.Left != nil {
				left.Push(root.Left)
			}
			root = root.Right
		} else {
			root, _ = left.Pop()
		}
	}
	for {
		node, ok := right.Pop()
		if !ok {
			break
		}
		res = append(res, node.Val)
	}
	return res
}


/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{}
	ret := [][]int{}
	q = append(q, root)
	for len(q) > 0 {
		// dequeue
		tmp := []int{}
		nodes := []*TreeNode{}
		for _, el := range q {
			tmp = append(tmp, el.Val)
			if el.Left != nil {
				nodes = append(nodes, el.Left)
			}
			if el.Right != nil {
				nodes = append(nodes, el.Right)
			}
		}
		ret = append(ret, tmp)
		q = nodes[:]
	}
	if len(ret) <= 1 {
		return ret
	}
	i := len(ret) - 1
	rev := [][]int{}
	for i >= 0 {
		rev = append(rev, ret[i])
		i--
	}
	return rev
}


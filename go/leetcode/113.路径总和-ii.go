/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (55.11%)
 * Likes:    88
 * Dislikes: 0
 * Total Accepted:    9.5K
 * Total Submissions: 17.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 * 
 * 
 * 返回:
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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

func helper(root *TreeNode, sum int)  {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	tmp := append([]int{}, path...)   // copy
	rest := sum - root.Val
	if root.Left == nil && root.Right == nil && rest == 0 {
		ret = append(ret, tmp)   // use copy, not ref
		return
	}
	helper(root.Left, rest)
	path = tmp    // remember reset path
	helper(root.Right, rest)
	if len(path) > 0 {
		path = path[0:len(path)-1]
	}
}

var ret [][]int
var path []int

func pathSum(root *TreeNode, sum int) [][]int {
	ret = [][]int{}
	path = []int{}
	helper(root, sum)
	return ret
}


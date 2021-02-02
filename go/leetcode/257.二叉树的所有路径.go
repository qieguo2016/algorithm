/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (66.43%)
 * Likes:    436
 * Dislikes: 0
 * Total Accepted:    93.5K
 * Total Submissions: 140.7K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * 输出: ["1->2->5", "1->3"]
 * 
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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

func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)
	helper(root, "", &ret)
	return ret
}

func helper(root *TreeNode, path string, ret *[]string) {
	if root == nil {
		return
	}
	var newPath string
	if path == "" {
		newPath = fmt.Sprintf("%d", root.Val)
	} else {
		newPath = fmt.Sprintf("%s->%d", path, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, newPath)
		return
	}
	if root.Left != nil {
		helper(root.Left, newPath, ret)
	}
	if root.Right != nil {
		helper(root.Right, newPath, ret)
	}
}
// @lc code=end


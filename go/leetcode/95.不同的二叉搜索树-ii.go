/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (59.00%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    11.2K
 * Total Submissions: 19K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
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

// 跟96题差不多，根结点从1到n循环，左节点[1,i-1]，右节点[i+1, n]
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}

func helper(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil} // 无子节点，用nil表示
	}
	ret := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		// 卡塔兰数，分两步组合，每一步组合是i和n-i的关系，总组合数=sum(i*n-i), i=[0,n]
		for _, ln := range left {
			for _, lr := range right {
				node := &TreeNode{
					Val:   i,
					Left:  ln,
					Right: lr,
				}
				ret = append(ret, node)
			}
		}
	}
	return ret
}

// @lc code=end

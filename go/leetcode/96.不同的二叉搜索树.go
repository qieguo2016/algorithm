/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (62.17%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    16.8K
 * Total Submissions: 27.1K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出: 5
 * 解释:
 * 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 * 
 */
func numTrees(n int) int {
    if n < 2 {
			return 1
		}
		if n == 2 {
			return 2
		}
		count := 0
		for i := 0; i < n; i++ {
			count += numTrees(i) * numTrees(n-i-1)
		}
		return count
}


/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 *
 * https://leetcode-cn.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (25.03%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    11.1K
 * Total Submissions: 44.3K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
 *
 * '?' 可以匹配任何单个字符。
 * '*' 可以匹配任意字符串（包括空字符串）。
 *
 *
 * 两个字符串完全匹配才算匹配成功。
 *
 * 说明:
 *
 *
 * s 可能为空，且只包含从 a-z 的小写字母。
 * p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
 *
 *
 * 示例 1:
 *
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串。
 *
 * 示例 2:
 *
 * 输入:
 * s = "aa"
 * p = "*"
 * 输出: true
 * 解释: '*' 可以匹配任意字符串。
 *
 *
 * 示例 3:
 *
 * 输入:
 * s = "cb"
 * p = "?a"
 * 输出: false
 * 解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
 *
 *
 * 示例 4:
 *
 * 输入:
 * s = "adceb"
 * p = "*a*b"
 * 输出: true
 * 解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
 *
 *
 * 示例 5:
 *
 * 输入:
 * s = "acdcb"
 * p = "a*c?b"
 * 输入: false
 *
 */

import (
	"strings"
)

const (
	w byte = '*'
	c byte = '?'
)

// 动态规划统一解法
// dp[i][j]表示s中前i个字符组成的子串和p中前j个字符组成的子串是否能匹配
// dp[m+1][n+1]是为了兼容空串的匹配
// 1. 初始状态 dp[0][0]为0；s[0]表示s长度为0，那么连续的*号也可以匹配
// 2. 转移方程 p[j]为*，*可以取空串或者任意串，那么dp[i][j-1]或dp[i-1][j]匹配即可
//    转移方程 p[j]不为*，dp[i-1][j-1]需要先匹配上，然后s[i]=p[j]或者p[j]=?
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		if p[j-1] == w {
			dp[0][j] = dp[0][j-1] // s为空串，连续*号为true
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == w { // 还是要从0开始遍历，又因为要取j-1的元素，所以用j-1来遍历
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-1] && (p[j-1] == c || p[j-1] == s[i-1])
			}
		}
	}
	return dp[m][n]
}


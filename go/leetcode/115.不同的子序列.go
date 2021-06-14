/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 *
 * https://leetcode-cn.com/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (53.43%)
 * Likes:    534
 * Dislikes: 0
 * Total Accepted:    52.5K
 * Total Submissions: 98.3K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
 *
 * 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE"
 * 的一个子序列，而 "AEC" 不是）
 *
 * 题目数据保证答案符合 32 位带符号整数范围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "rabbbit", t = "rabbit"
 * 输出：3
 * 解释：
 * 如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
 * (上箭头符号 ^ 表示选取的字母)
 * rabbbit
 * ^^^^ ^^
 * rabbbit
 * ^^ ^^^^
 * rabbbit
 * ^^^ ^^^
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "babgbag", t = "bag"
 * 输出：5
 * 解释：
 * 如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
 * (上箭头符号 ^ 表示选取的字母)
 * babgbag
 * ^^ ^
 * babgbag
 * ^^    ^
 * babgbag
 * ^    ^^
 * babgbag
 * ⁠ ^  ^^
 * babgbag
 * ⁠   ^^^
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 和 t 由英文字母组成
 *
 *
 */

// @lc code=start
func numDistinct(s string, t string) int {
	// 匹配的情况枚举：
	// 1. s[i]==t[j]，用或者不用两种
	//    用，再看s[i+1:]和t[j+1:]
	//    不用，再看s[i+1:]和t[j:]
	// 2. s[i]!=t[j]，看s[i+1:]和t[j:]
	// dp定义：dp[i,j]表示s[i:]和t[j:]能匹配的数量，结果就是求d[0,0]
	// 边界，dp[x,lt]=1，空串是任意字符串的子串，dp[ls,x]=0，任意字符串都不是空串的子串
	// 状态转移参考上面分析可得
	// dp[i,j] = dp[i+1,j+1] + dp[i+1,j]  (s[i]==t[j])
	// dp[i,j] = dp[i+1,j]  (s[i]!=t[j])
	// tc = O(ls*lt)  sc=(ls*lt)

	// 快速退出
	ls, lt := len(s), len(t)
	if ls <= 0 {
		return 0
	}
	if lt <= 0 {
		return 1
	}

	// 边界
	dp := make([][]int, ls+1)
	for i := 0; i < ls+1; i++ {
		dp[i] = make([]int, lt+1)
		dp[i][lt] = 1
	}
	for j := 0; j < lt+1; j++ {
		dp[ls][j] = 0
	}
	dp[ls][lt] = 1 // 空串是空串的子串，上一个for循环覆盖了

	for i := ls - 1; i >= 0; i-- {
		for j := lt - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

// @lc code=end


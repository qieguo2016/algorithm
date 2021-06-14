/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 *
 * https://leetcode-cn.com/problems/interleaving-string/description/
 *
 * algorithms
 * Medium (45.74%)
 * Likes:    464
 * Dislikes: 0
 * Total Accepted:    50.2K
 * Total Submissions: 109.7K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
 *
 * 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
 *
 *
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tm
 * |n - m|
 * 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 +
 * ...
 *
 *
 * 提示：a + b 意味着字符串 a 和 b 连接。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：s1 = "", s2 = "", s3 = ""
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 0
 * s1、s2、和 s3 都由小写英文字母组成
 *
 *
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	// 必要条件：len(s1)+len(s2)=len(s3)
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	if l3 <= 0 {
		return true
	}
	// 动态规划，dp[i,j]表示s3[:i+j]能否由s1[:i]和s2[:j]交错而成
	// 那么s3[i+j]这个字符要么是s1[i]，要么是s2[j]，否则不可能
	// 1. s3[i+j]=s1[i]，那么就要看s3[i+j-1]能否由s1[i-1]和s2[j]交错而成
	// 2. s3[i+j]=s2[j]，那么就要看s3[i+j-1]能否由s1[i]和s2[j-1]交错而成
	// dp[i,j] = (s3[i+j-1]=s1[i-1] && dp[i-1, j]) || (s3[i+j-1]=s2[j-1] && dp[i, j-1])
	// i+j-1和i-1,j-1都是因为字符串数组从0开始
	// dp := make([][]bool, l1+1) // 0号元素用来递推
	// for i := 0; i <= l1; i++ {
	// 	dp[i] = make([]bool, l2+1) // 0号元素用来递推
	// }
	// dp[0][0] = true
	// for i := 0; i <= l1; i++ {
	// 	for j := 0; j <= l2; j++ {
	// 		k := i + j - 1
	// 		if i > 0 && s3[k] == s1[i-1] {
	// 			dp[i][j] = dp[i-1][j]
	// 		}
	// 		if !dp[i][j] && j > 0 && s3[k] == s2[j-1] { // dp[i][j]确定可以被分解就可以递推了
	// 			dp[i][j] = dp[i][j-1]
	// 		}
	// 	}
	// }
	// return dp[l1][l2]

	// 上面的时间空间复杂度都是O(l1*l2)
	// dp[i,j] = (s3[i+j-1]=s1[i-1] && dp[i-1, j]) || (s3[i+j-1]=s2[j-1] && dp[i, j-1])
	// 简化一点就是 dp[i,j] = (x && dp[i-1, j]) || (dp[i, j-1] && y)
	// 每一次外层循环都是更新列，也就是列可以看成多行的叠加，换成列就是 dp[j] = (x && dp[j]) || (dp[j-1] && y)
	dp := make([]bool, l2+1) // 0号元素用来递推
	dp[0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			k := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s3[k] == s1[i-1]
			}
			if !dp[j] && j > 0 {
				dp[j] = dp[j-1] && s3[k] == s2[j-1]
			}
		}
	}
	return dp[l2]

}

// @lc code=end


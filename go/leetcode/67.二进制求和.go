/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (49.08%)
 * Likes:    207
 * Dislikes: 0
 * Total Accepted:    28.2K
 * Total Submissions: 57.5K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */

// 注意进位
// package leetcode

// @lc code=start
func addBinary(a string, b string) string {
	ret := ""
	carry := 0
	byteA := []byte(a)
	byteB := []byte(b)
	m := len(a) - 1
	n := len(b) - 1
	var zero byte = '0'
	for m >= 0 || n >= 0 {
		ca := 0
		if m >= 0 {
			ca = int(byteA[m] - zero)
			m--
		}
		cb := 0
		if n >= 0 {
			cb = int(byteB[n] - zero)
			n--
		}
		sum := ca + cb + carry
		ret = string(zero+byte(sum%2)) + ret
		carry = sum / 2
	}
	if carry > 0 {
		ret = "1" + ret
	}
	return ret
}

// @lc code=end

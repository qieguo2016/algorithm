/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (24.38%)
 * Total Accepted:    38.2K
 * Total Submissions: 155.7K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
func findPalindrome(s string, left int, right int) string {
	j := left
	k := right
	l := len(s)
	ret := s[j : k+1]
	for {
		j--
		k++
		if j < 0 || k > l-1 {
			break
		}
		if s[j] != s[k] {
			break
		}
		ret = s[j : k+1]
	}
	return string(ret)
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	target := s[:1]
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			p := findPalindrome(s, i, i+1)
			if len(target) < len(p) {
				target = p
			}
		}
		if i-1 >= 0 && i+1 < l && s[i-1] == s[i+1] {
			p := findPalindrome(s, i-1, i+1)
			if len(target) < len(p) {
				target = p
			}
		}
	}
	return target
}


/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (31.45%)
 * Total Accepted:    47.4K
 * Total Submissions: 150.5K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {
	ret := []byte{}
	i := 0
	if len(strs) == 0 {
		return ""
	}
	for {
		var c byte
		for _, str := range strs {
			if i > len(str)-1 {
				return string(ret)
			}
			if c == 0 {
				c = str[i]
				continue
			}
			if c != str[i] {
				return string(ret)
			}
		}
		ret = append(ret, c)
		i++
	}
	return ""
}

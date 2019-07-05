/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
import (
)
func lengthOfLastWord(s string) int {
	r := len(s) - 1
	res := 0
	for r >=0 && s[r] == " "[0] {
		r--
	}
	for r >= 0 && s[r] != " "[0] {
		r--
		res++
	}
	return res
}


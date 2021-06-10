/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (26.13%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    7.3K
 * Total Submissions: 27.8K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 * 
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 * 
 * 
 * 示例 2：
 * 
 * 输入：
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * 输出：[]
 * 
 * 
 */

// @lc code=start
// 使用hash记录words及对应次数，按子串组合长度分段遍历目标字符串，判断是否可以由words组成
func findSubstring(s string, words []string) []int {
	ret := []int{}
  	if len(words) <= 0 {
		return ret
	}
	m := len(words)
	n := len(words[0])
	step := m*n
	start := 0
	wm := map[string]int{}
	for _, word := range words {
		if v, ok := wm[word]; ok {
			wm[word] = v + 1
		} else {
			wm[word] = 1
		}
	}
	for start + step <= len(s) {
		equal := 0
		r := map[string]int{}
		for i := 0; i < m; i++ {
			st := start+i*n
			word := s[st: st+n]
			if vt, ok := wm[word]; ok {
				if vc, ok := r[word]; ok {
					if vc < vt {   // duplicate word
						equal++
					}
					r[word] = vc + 1
					} else {   // first time
					r[word] = 1
					equal++
				}
			}
		}
		if equal == m {
			ret = append(ret, start)
		}
		start++
	}
	return ret
}
// @lc code=end


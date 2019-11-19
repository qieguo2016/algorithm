/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (58.35%)
 * Likes:    192
 * Dislikes: 0
 * Total Accepted:    31.3K
 * Total Submissions: 53.4K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

import (
	"sort"
)

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	cm := map[[26]int]int{}
	ret := [][]string{}
	var a rune = 'a'
	for _, str := range strs {
		ca := [26]int{}
		for _, c := range str {
			ca[c-a] += 1
		}
		if i, ok := cm[ca]; ok {
			ret[i] = append(ret[i], str)
		} else {
			ret = append(ret, []string{str})
			cm[ca] = len(ret) - 1
		}
	}
	return ret
}
// @lc code=end


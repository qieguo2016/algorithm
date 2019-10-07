/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (49.81%)
 * Likes:    361
 * Dislikes: 0
 * Total Accepted:    29.9K
 * Total Submissions: 60.1K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 * 
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 * 
 */

// @lc code=start
var mm = [][]string{
	[]string{"a","b", "c"},  // 2
	[]string{"d", "e", "f"},
	[]string{"g", "h", "i"},
	[]string{"j", "k", "l"},
	[]string{"m", "n", "o"},
	[]string{"p", "q", "r", "s"},
	[]string{"t", "u", "v"},
	[]string{"w", "x", "y", "z"},  // 9
}
const two rune = '2'

func dfs(cs *[]rune, level int, out string, res *[]string) {
	if level >= len(*cs) {
		*res = append(*res, out)  // 复制值
		return 
	}
	i := int((*cs)[level]-two)
	if i < 0 || i >= len(mm) {
		return
	}
	str := mm[i]
	for _, s := range str {
		dfs(cs, level+1, out+s, res)
	}
}

func letterCombinations(digits string) []string {
	ret := []string{}
	n := len(digits)
	if n <= 0 {
		return ret
	}
	cs := []rune(digits)
	dfs(&cs, 0, "", &ret)
	return ret
}
// @lc code=end


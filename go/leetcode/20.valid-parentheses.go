import "fmt"

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.00%)
 * Total Accepted:    42.9K
 * Total Submissions: 118.9K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
func isValid(s string) bool {
	if s == "" {
		return true
	}
	// '('，')'，'{'，'}'，'['，']'
	left := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	right := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := []string{}
	for _, cr := range s {
		c := fmt.Sprintf("%c", cr)
		if _, exist := left[c]; exist {
			stack = append(stack, c)
			continue
		}
		w, exist := right[c]
		si := len(stack) - 1
		if !exist || si < 0 || w != stack[si] {
			return false
		}
		stack = stack[:si]
	}
	return len(stack) == 0
}

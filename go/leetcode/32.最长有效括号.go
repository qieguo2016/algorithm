/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (27.02%)
 * Likes:    230
 * Dislikes: 0
 * Total Accepted:    11.4K
 * Total Submissions: 42K
 * Testcase Example:  '"(()"'
 *
 * 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
 * 
 * 示例 1:
 * 
 * 输入: "(()"
 * 输出: 2
 * 解释: 最长有效括号子串为 "()"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ")()())"
 * 输出: 4
 * 解释: 最长有效括号子串为 "()()"
 * 
 * 
 */

// 括号匹配一般使用栈来做，而且在括号匹配中，只有左括号才能入栈，那么栈中元素的值已经不重要了，只有有值就说明有左括号
// 回到这个题目，则可以将左括号的序号入栈，匹配时根据序号差确定子串长度
// 出栈之后栈空了，所以需要一个变量存括号起始位置：()()、start=0
// 出栈后栈顶的元素是未匹配到的左括号位置，也就是有效子串的起始位置
func longestValidParentheses(s string) int {
	// ()((() 2
	stack := []int{}
	left := "("[0]
	max := 0
	start := 0
	for i := range s {
		if s[i] == left {
			stack = append(stack, i)   // 左括号直接入栈
		} else if len(stack) <= 0 {  // 空栈遇到右括号，有效括号在这里截断，下标更新
			start = i + 1
		} else {  // 非空，右括号能匹配到左括号
			stack = stack[:len(stack)-1]
			var tmp int
			if len(stack) <= 0 {  // 取出之后栈空了，那就从有效下标开始
				tmp = i - start + 1
			} else {
				tmp = i - stack[len(stack)-1]  // 对应(()()这种情况，栈顶元素是有效子串的起始位置
			}
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}


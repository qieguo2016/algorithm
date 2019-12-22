/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (34.08%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    8.5K
 * Total Submissions: 25K
 * Testcase Example:  '"3+2*2"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
 *
 * 示例 1:
 *
 * 输入: "3+2*2"
 * 输出: 7
 *
 *
 * 示例 2:
 *
 * 输入: " 3/2 "
 * 输出: 1
 *
 * 示例 3:
 *
 * 输入: " 3+5 / 2 "
 * 输出: 5
 *
 *
 * 说明：
 *
 *
 * 你可以假设所给定的表达式都是有效的。
 * 请不要使用内置的库函数 eval。
 *
 *
 */

// package leetcodex
// 用栈实现，遍历字符串，操作符为乘除的时候要先计算后入栈
// 1. 都是两个数的运算，要依次得到左右数和运算符，左数直接入栈，操作符和右数用额外空间存储
// 2. 遍历到最后一个右运算数的时候要进入计算环节

// @lc code=start
func calculate(s string) int {
	stack := []int{}
	var zero rune = '0'
	var nine rune = '9'
	var plus rune = '+'
	var minus rune = '-'
	var multiply rune = '*'
	var divide rune = '/'

	num := 0
	op := plus // 记录上一次操作符
	for i, c := range s {
		if c >= zero && c <= nine {
			num = num*10 + int(c-zero) // 3,5,2
		}
		if c == plus || c == minus || c == multiply || c == divide || i == len(s)-1 {
			switch op {
			case plus:
				stack = append(stack, num) // 3,5
			case minus:
				stack = append(stack, -num)
			case multiply:
				stack[len(stack)-1] *= num
			case divide:
				stack[len(stack)-1] /= num // 5/2
			}
			num = 0
			op = c // + /
		}
	}
	res := 0
	for i := range stack {
		res += stack[i]
	}
	return res
}

// @lc code=end

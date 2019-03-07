/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (31.16%)
 * Total Accepted:    75.3K
 * Total Submissions: 241.4K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
 *
 */
func reverse(x int) int {
	ret := 0
	int_max := 1<<31 - 1 // 2147483647
	int_min := ^int_max  // -2147483648
	for true {
		if x == 0 {
			break
		}
		// 提前判断ret是否溢出时，注意x本身就在int范围内，只要保证后面+x%10不溢出就可以了
		// 如果最后一步ret=214748364，根据x范围那么最后一位只能是1，也就是x的第一位是1，x=1463847412，ret=2147483641，不会溢出
		if ret > int_max/10 || ret < int_min/10 {
			return 0
		}
		ret = ret*10 + x%10
		x /= 10
	}
	return ret
}

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (18.27%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    13.2K
 * Total Submissions: 72.2K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 * 
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 * 
 * 示例 1:
 * 
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 
 * 示例 2:
 * 
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 
 * 说明:
 * 
 * 
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 * 
 * 
 */

// 最简单的方式是不断减被除数，可减的次数就是商；可以用位移来将被除数乘2，从而加速运算
func divide(dividend int, divisor int) int {
	int_max := 1<<31 - 1 // 2147483647
	int_min := ^int_max  // -2147483648
	if divisor == 0 {
		return int_max
	}

	flag := -1
	if (dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0) {
		flag = 1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	s := uint64(dividend)
	p := uint64(divisor)
	i := 1
	ret := 0
	for s >= p {
		tmp := p << 1  // *2
		if s > tmp {  // 10 6
			p = tmp   // 6
			i *= 2  // 2
		} else {  // 10 12
			ret += i   // 2, 3
			s -= p  // 10 - 6 = 4 - 3 = 1
			p = uint64(divisor) // 3
			i = 1
		}
	}
	ret = flag*ret
	if ret > int_max {
		return int_max
	}
	if ret < int_min {
		return int_max
	}
	return ret
}


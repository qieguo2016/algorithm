/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (40.37%)
 * Likes:    214
 * Dislikes: 0
 * Total Accepted:    29K
 * Total Submissions: 71.5K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 * 
 * 示例 1:
 * 
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 * 
 * 示例 2:
 * 
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 * 
 * 说明：
 * 
 * 
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 * 
 * 
 */

// @lc code=start
import (
	"strconv"
)

const zero byte = '0'
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m := len(num1)
	n := len(num2)
	digit := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {  // 数据的低位存在数组的高位
			tmp := int(num1[i] - zero) * int(num2[j] - zero) + digit[i+j+1]  // 从0开始所以+1
			digit[i+j+1] = tmp % 10  // 数组的高位存数据的低位
			digit[i+j] += tmp / 10   // 数组的低位存数据的高位
		}
	}
	k := 0
	for ; k < len(digit) && digit[k] == 0; k++ {
		// 清理高位的0
	}
	res := ""
	for ; k < len(digit); k++ {
		res += strconv.Itoa(digit[k])
	}
	return res
}
// @lc code=end

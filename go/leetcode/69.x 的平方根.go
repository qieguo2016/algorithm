/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (36.15%)
 * Likes:    176
 * Dislikes: 0
 * Total Accepted:    40.1K
 * Total Submissions: 111K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

import (
	"math"
)
// 牛顿迭代法，x1=(x0+n/x0)/2
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var res float64 = 1
	var pre float64 = 0
	for math.Abs(res - pre) > 1e-6 {
		pre = res
		res = (res + float64(x) / res) / 2
	}
	return int(res)
}


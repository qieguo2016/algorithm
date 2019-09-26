/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (50.14%)
 * Likes:    442
 * Dislikes: 0
 * Total Accepted:    54.1K
 * Total Submissions: 107.8K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 * 
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
 * 
 * 注意你不能在买入股票前卖出股票。
 * 
 * 示例 1:
 * 
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 
 */

// 假设买点在第一天，后面发现如果有更低的，就更新买点。
// 在发现更低买点的时候，是不可能在这个时候卖出的，所以最大利润是昨天的利润
// 遍历下来，不断更新买点和比较最大利润，就可以得到结果
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := prices[0]
	res := 0
  for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		}
		tmp := prices[i] - buy
		if tmp > res {
			res = tmp
		}
	}
	return res
}


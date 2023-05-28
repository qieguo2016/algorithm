/*
 * @lc app=leetcode.cn id=123 lang=cpp
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (58.68%)
 * Likes:    1413
 * Dislikes: 0
 * Total Accepted:    256.8K
 * Total Submissions: 437.7K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：prices = [3,3,5,0,0,3,1,4]
 * 输出：6
 * 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 =
 * 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。 随后，在第 7 天（股票价格 =
 * 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 =
 * 4-1 = 3 。
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 =
 * 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。   注意你不能在第 1 天和第 2
 * 天接连购买股票，之后再将它们卖出。  
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 *
 * 示例 3：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
 *
 * 示例 4：
 *
 *
 * 输入：prices = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 某一天的的状态有几种：
  // 1 未操作，2 买入一次，3买卖一次，4买卖一次继续买，5，完成两次买卖
  // 1无利润不管，设其他几种状态下的最大利润分别为buy1,sell1,buy2,sell2
  // 分别更新每天这几种状态的最大值，最终输出0，sell1, sell2中的最大值
  int maxProfit(vector<int> &prices) {
    int n = prices.size();
    if (n <= 1) {
      return 0;
    }
    int buy1 = -prices[0], sell1 = 0;
    int buy2 = -prices[0], sell2 = 0;

    for (int i = 1; i < n; i++) {
      buy1 = max(buy1, -prices[i]); // 不买入buy1，买入时利润为-prices[i]
      sell1 = max(sell1, buy1 + prices[i]); // 卖出为buy1 + prices[i]
      buy2 = max(buy2, sell1 - prices[i]);  // 买入则为sell1 - prices[i]
      sell2 = max(sell2, buy2 + prices[i]); // 卖出则为buy2 + prices[i]
    }
    return sell2; // sell1也会转移到sell2
  }
};
// @lc code=end

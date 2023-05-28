/*
 * @lc app=leetcode.cn id=279 lang=cpp
 *
 * [279] 完全平方数
 *
 * https://leetcode.cn/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (66.14%)
 * Likes:    1703
 * Dislikes: 0
 * Total Accepted:    403.8K
 * Total Submissions: 610.5K
 * Testcase Example:  '12'
 *
 * 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
 *
 * 完全平方数
 * 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9
 * 和 16 都是完全平方数，而 3 和 11 不是。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12
 * 输出：3
 * 解释：12 = 4 + 4 + 4
 *
 * 示例 2：
 *
 *
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */

#include <climits>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 完全背包问题，可选项是1~根号n内的平方数
  int numSquares(int n) {
    if (n <= 3) {
      return n;
    }
    vector<int> dp(n + 1);
    for (int i = 1; i < n + 1; i++) {
      int mini = INT_MAX;
      for (int j = 1; j * j <= i; j++) {
        // dp(i)的最后一步可以是1~根号i内的所有平方数，所以步数等于min(dp[i - j * j]) + 1
        mini = min(mini, dp[i - j * j]);
      }
      dp[i] = mini + 1;
    }
    return dp[n];
  }
};
// @lc code=end

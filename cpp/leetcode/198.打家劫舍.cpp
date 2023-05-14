/*
 * @lc app=leetcode.cn id=198 lang=cpp
 *
 * [198] 打家劫舍
 *
 * https://leetcode.cn/problems/house-robber/description/
 *
 * algorithms
 * Medium (54.31%)
 * Likes:    2560
 * Dislikes: 0
 * Total Accepted:    742.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下
 * ，一夜之内能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 *
 *
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋
 * (金额 = 1)。 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
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

// @lc code=start
class Solution {
public:
  // 这种dp题目，一般是分析某个位置要或丢两种选择怎么选，输出结果有两种模式：
  // 一种是维护一个全局最大值，每次循环比较更新;
  // 另外一种是直接求dp[n]或者dp[0]的值
  // dp[i]标识i个数字内的最大值，因为数字都是正数，只要不触发报警机制都可以要，那要第i个数的前提是第i-1个没选
  // 因此 dp[i] = max(dp[i-2]+nums[i], dp[i-1])
  int rob(std::vector<int> &nums) {
    if (nums.empty()) {
      return 0;
    }
    if (nums.size() < 2) {
      return nums[0];
    }
    std::vector<int> dp(nums.size());
    dp[0] = nums[0];
    dp[1] = std::max(nums[0], nums[1]);
    for (size_t i = 2; i < nums.size(); i++) {
      dp[i] = std::max(dp[i - 2] + nums[i], dp[i - 1]);
    }
    return dp[nums.size() - 1];
  }
};
// @lc code=end

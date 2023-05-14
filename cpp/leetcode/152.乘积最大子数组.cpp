/*
 * @lc app=leetcode.cn id=152 lang=cpp
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode.cn/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (43.11%)
 * Likes:    2012
 * Dislikes: 0
 * Total Accepted:    356.5K
 * Total Submissions: 826.9K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组
 * nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 * 测试用例的答案是一个 32-位 整数。
 *
 * 子数组 是数组的连续子序列。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
 *
 *
 */
#include <algorithm>
#include <vector>
// @lc code=start
class Solution {
public:
  // 与最大和类似，区别是负数后面还可以再乘负数得到正数，而正数也可以乘负数得到正数，而最大和最需要关注正数即可。
  // 因此可以用两个变量分别存正数最大值和负数最小值，每次循环判断当前值去更新这两个值，同时更新全局最大值
  int maxProduct(std::vector<int> &nums) {
    if (nums.empty()) {
      return 0;
    }
    int res = nums[0], maximum = nums[0], minimum = nums[0];
    for (size_t i = 1; i < nums.size(); i++) {
      if (nums[i] > 0) {
        maximum = std::max(maximum * nums[i], nums[i]);
        minimum = std::min(minimum * nums[i], nums[i]);
      } else {
        int tmp = maximum;
        maximum = std::max(minimum * nums[i], nums[i]);
        minimum = std::min(tmp * nums[i], nums[i]);
      }
      res = std::max(maximum, res);
    }
    return res;
  }
};
// @lc code=end

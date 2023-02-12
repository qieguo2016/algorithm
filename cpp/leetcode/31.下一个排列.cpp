/*
 * @lc app=leetcode.cn id=31 lang=cpp
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (37.10%)
 * Likes:    1237
 * Dislikes: 0
 * Total Accepted:    188.9K
 * Total Submissions: 507.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取 下一个排列
 * 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 *
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 *
 * 必须 原地 修改，只允许使用额外常数空间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[1,2,3]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,1,5]
 * 输出：[1,5,1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [1]
 * 输出：[1]
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

#include <algorithm>
#include <vector>
// @lc code=start
class Solution {
public:
  void nextPermutation(std::vector<int> &nums) {
    // [1,4,9,8,3,2] => [1,8,2,3,4,9] => swap48，revert9432
    if (nums.size() <= 1) {
      return;
    }
    int i = nums.size() - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) {
      i--;
    }
    if (i >= 0) { // found
      int j = nums.size() - 1;
      while (nums[j] <= nums[i]) {
        j--;
      }
      std::swap(nums[i], nums[j]);
    }
    std::reverse(nums.begin() + i + 1, nums.end());
  }
};
// @lc code=end

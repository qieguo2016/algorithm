/*
 * @lc app=leetcode.cn id=41 lang=cpp
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (41.26%)
 * Likes:    1104
 * Dislikes: 0
 * Total Accepted:    136.6K
 * Total Submissions: 331.2K
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
 * 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,0]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,4,-1,1]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,8,9,11,12]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 *
 *
 */

#include <vector>

// @lc code=start
class Solution {
public:
  int firstMissingPositive(std::vector<int> &nums) {
    // [3,4,0,1] => 2,   -1,1,4,3
    for (size_t i = 0; i < nums.size(); i++) {
      while (nums[i] > 0 && nums[i] < nums.size() &&
             nums[i] != nums[nums[i] - 1]) { // loop while equal
        std::swap(nums[i], nums[nums[i] - 1]);
      }
    }

    for (size_t i = 0; i < nums.size(); i++) {
      if (nums[i] != i + 1) {
        return i + 1;
      }
    }
    return nums.size() + 1;
  }
};
// @lc code=end

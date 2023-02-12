/*
 * @lc app=leetcode.cn id=34 lang=cpp
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.52%)
 * Likes:    1123
 * Dislikes: 0
 * Total Accepted:    295.7K
 * Total Submissions: 695.1K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值
 * target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 *
 * 进阶：
 *
 *
 * 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 *
 * 示例 3：
 *
 *
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^9 
 * nums 是一个非递减数组
 * -10^9 
 *
 *
 */

#include <vector>

// @lc code=start
class Solution {
public:
  std::vector<int> searchRange(std::vector<int> &nums, int target) {
    std::vector<int> ret = {-1, -1};
    if (nums.empty()) {
      return ret;
    }
    int left = 0, right = nums.size();
    while (left < right) {
      int mid = left + (right - left) / 2;
      if (target <= nums[mid]) {
        right = mid;
      } else {
        left = mid + 1;
      }
    }
    if (left >= nums.size() || nums[left] != target) {
      return ret;
    }
    ret[0] = left;
    left = 0, right = nums.size();
    while (left < right) {
      int mid = left + (right - left) / 2;
      if (target < nums[mid]) {
        right = mid;
      } else {
        left = mid + 1;
      }
    }
    ret[1] = right - 1; // 第一轮已经找到
    return ret;
  }
};
// @lc code=end

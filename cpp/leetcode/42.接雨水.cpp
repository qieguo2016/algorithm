/*
 * @lc app=leetcode.cn id=42 lang=cpp
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (56.05%)
 * Likes:    2405
 * Dislikes: 0
 * Total Accepted:    257.1K
 * Total Submissions: 458.7K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1
 * 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1]
 * 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 0
 * 0
 *
 *
 */

#include <vector>

// @lc code=start
class Solution {
public:
  int trap(std::vector<int> &height) {
    int left = 0, right = height.size() - 1;
    int leftMax = 0, rightMax = 0;
    int ret = 0;
    while (left < right) {
      if (height[left] < height[right]) {
        if (height[left] < leftMax) {
          ret += leftMax - height[left];
        } else {
          leftMax = height[left];
        }
        left++;
      } else {
        if (height[right] < rightMax) {
          ret += rightMax - height[right];
        } else {
          rightMax = height[right];
        }
        right--;
      }
    }
    return ret;
  }
};
// @lc code=end

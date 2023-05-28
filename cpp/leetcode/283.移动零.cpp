/*
 * @lc app=leetcode.cn id=283 lang=cpp
 *
 * [283] 移动零
 *
 * https://leetcode.cn/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.80%)
 * Likes:    2012
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.7M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0
 * 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [0]
 * 输出: [0]
 *
 *
 *
 * 提示:
 *
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能尽量减少完成的操作次数吗？
 *
 */

#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 经典双指针
  void moveZeroes(vector<int> &nums) {
    int n = nums.size();
    if (n <= 1) {
      return;
    }
    int l = 0, r = 0; // l左侧全非0，lr之间全是0
    while (r < n) {
      // [0,1,0,3,12]
      if (nums[r] != 0) {
        swap(nums[l++], nums[r++]);
      } else {
        r++;
      }
    }
  }
};
// @lc code=end

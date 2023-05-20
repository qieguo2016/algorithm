/*
 * @lc app=leetcode.cn id=209 lang=cpp
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (46.90%)
 * Likes:    1706
 * Dislikes: 0
 * Total Accepted:    541.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。
 *
 * 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ...,
 * numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：target = 4, nums = [1,4,4]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n))
 * 时间复杂度的解法。
 *
 *
 */
#include <algorithm>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 双指针遍历，先右指针向前直到左右指针间和大于目标，再左指针向前直到和小于目标，不断更新最小长度
  int minSubArrayLen(int target, vector<int> &nums) {
    if (nums.empty() || target == 0) {
      return 0;
    }
    int l = 0, r = 0, s = 0, ret = nums.size();
    while (r < nums.size()) {
      while (r < nums.size() && s < target) {
        s += nums[r++];
      }
      if (s < target) {
        return l == 0 ? 0 : ret;
      }
      ret = min(ret, r - l);
      while (l < r && s >= target) {
        s -= nums[l++];
      }
      ret = min(ret, r - l + 1);
    }
    return ret;
  }
};
// @lc code=end

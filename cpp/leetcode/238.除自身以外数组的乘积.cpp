/*
 * @lc app=leetcode.cn id=238 lang=cpp
 *
 * [238] 除自身以外数组的乘积
 *
 * https://leetcode.cn/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (74.80%)
 * Likes:    1457
 * Dislikes: 0
 * Total Accepted:    270.8K
 * Total Submissions: 362K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个整数数组 nums，返回
 * 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
 *
 * 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位
 * 整数范围内。
 *
 * 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,3,4]
 * 输出: [24,12,8,6]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [-1,1,0,-3,3]
 * 输出: [0,0,9,0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内
 *
 *
 *
 *
 * 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（
 * 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
 *
 */

#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 左右两个数组，分别代表当前元素左侧乘积和右侧乘积，两者想乘即为当前位置结果
  vector<int> productExceptSelf(vector<int> &nums) {
    int n = nums.size();
    if (nums.empty()) {
      return vector<int>();
    }
    vector<int> left(n), right(n);
    left[0] = 1, right[n - 1] = 1;
    for (int i = 1; i < n; i++) {
      left[i] = left[i - 1] * nums[i - 1];
    }
    for (int i = n - 2; i >= 0; i--) {
      right[i] = right[i + 1] * nums[i + 1];
    }

    vector<int> ret(n);
    for (int i = 0; i < n; i++) {
      ret[i] = left[i] * right[i];
    }
    return ret;
  }
};
// @lc code=end

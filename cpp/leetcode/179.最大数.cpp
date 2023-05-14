/*
 * @lc app=leetcode.cn id=179 lang=cpp
 *
 * [179] 最大数
 *
 * https://leetcode.cn/problems/largest-number/description/
 *
 * algorithms
 * Medium (41.13%)
 * Likes:    1125
 * Dislikes: 0
 * Total Accepted:    197.3K
 * Total Submissions: 479.6K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数
 * nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
 *
 * 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,2]
 * 输出："210"
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,30,34,5,9]
 * 输出："9534330"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 10^9
 *
 *
 */
#include <algorithm>
#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  string largestNumber(vector<int> &nums) {
    if (nums.empty()) {
      return "";
    }
    vector<string> arr(nums.size());
    for (size_t i = 0; i < nums.size(); i++) {
      arr[i] = to_string(nums[i]);
    }
    sort(arr.begin(), arr.end(), [](const auto &a, const auto &b) {
      return a + b > b + a; // 3, 34 and 3, 32
    });
    if (arr[0] == "0") {
      return "0";
    }
    string ret;
    for (const auto &s : arr) {
      ret += s;
    }
    return ret;
  }
};
// @lc code=end

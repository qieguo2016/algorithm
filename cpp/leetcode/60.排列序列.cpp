/*
 * @lc app=leetcode.cn id=60 lang=cpp
 *
 * [60] 排列序列
 *
 * https://leetcode.cn/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (53.42%)
 * Likes:    754
 * Dislikes: 0
 * Total Accepted:    123.3K
 * Total Submissions: 230.8K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3, k = 3
 * 输出："213"
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 4, k = 9
 * 输出："2314"
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 3, k = 1
 * 输出："123"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

#include <string>
#include <vector>

// @lc code=start
class Solution {
public:
  std::string getPermutation(int n, int k) {
    std::vector<int> factorial(n);
    factorial[0] = 1;
    for (size_t i = 1; i < n; i++) {
      factorial[i] = factorial[i - 1] * i;
    }
    // [1,1,2,6]
    std::string ret;
    int j;
    k--; // 从0开始计数
    std::vector<int> nums(n);
    for (size_t i = 0; i < nums.size(); i++) {
      nums[i] = i + 1;
    }
    for (int i = n - 1; i >= 0; i--) {
      j = k / factorial[i]; // 2
      ret += (nums[j] + '0');
      nums.erase(nums.begin() + j);
      k %= factorial[i];
    }
    return ret;
  }
};
// @lc code=end

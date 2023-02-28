/*
 * @lc app=leetcode.cn id=77 lang=cpp
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (76.85%)
 * Likes:    602
 * Dislikes: 0
 * Total Accepted:    172.3K
 * Total Submissions: 224.2K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */

#include <vector>

// @lc code=start
class Solution {
private:
  void helper(std::vector<std::vector<int>> &res, std::vector<int> &out, int n,
              int k, int start) {
    if (k == 0) {
      res.emplace_back(out);
      return;
    }
    for (int i = start; i <= n; i++) {
      out.emplace_back(i);
      helper(res, out, n, k - 1, i + 1);
      out.pop_back();
    }
  }

public:
  std::vector<std::vector<int>> combine(int n, int k) {
    std::vector<std::vector<int>> res;
    std::vector<int> out;
    helper(res, out, n, k, 1);
    return res;
  }
};
// @lc code=end

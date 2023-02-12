/*
 * @lc app=leetcode.cn id=22 lang=cpp
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.56%)
 * Likes:    3031
 * Dislikes: 0
 * Total Accepted:    639.4K
 * Total Submissions: 824.2K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且
 * 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

#include <string>
#include <vector>

// @lc code=start
class Solution {
private:
  void helper(int left, int right, std::string &cur,
              std::vector<std::string> &result) {
    if (left > right) {
      return;
    }
    if (left == 0 && right == 0) {
      result.push_back(cur);
      return;
    }
    if (left > 0) {
      cur.push_back('(');
      helper(left - 1, right, cur, result);
      cur.pop_back();
    }
    if (right > 0) {
      cur.push_back(')');
      helper(left, right - 1, cur, result);
      cur.pop_back();
    }
  }

public:
  std::vector<std::string> generateParenthesis(int n) {
    std::vector<std::string> result;
    std::string cur;
    helper(n, n, cur, result);
    return result;
  }
};
// @lc code=end

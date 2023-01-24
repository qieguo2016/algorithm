/*
 * @lc app=leetcode.cn id=5 lang=cpp
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (34.41%)
 * Likes:    3723
 * Dislikes: 0
 * Total Accepted:    608.5K
 * Total Submissions: 1.8M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "ac"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由数字和英文字母（大写和/或小写）组成
 *
 *
 */

// @lc code=start
#include <iostream>
#include <string>
#include <utility>
class Solution {
public:
  std::string longestPalindrome(std::string s) {
    int start = 0, end = 0;
    for (size_t i = 0; i < s.size(); i++) {
      auto p1 = findPalindrome(s, i, i);
      auto p2 = findPalindrome(s, i, i + 1);
      if (p1.second - p1.first > end - start) {
        start = p1.first;
        end = p1.second;
      }
      if (p2.second - p2.first > end - start) {
        start = p2.first;
        end = p2.second;
      }
    }
    return s.substr(start, end- start + 1);
  }

  std::pair<int, int> findPalindrome(const std::string &s, int left,
                                     int right) {
    while (left >= 0 && right < s.size() && s[left] == s[right]) {
      left--;
      right++;
    }
    return {left + 1, right - 1};
  }
};
// @lc code=end

void print_ret(std::string ret) {
  std::cout << "ret: ";
  std::cout << ret << " ";
  std::cout << std::endl;
}

int main(int argc, const char **argv) {
  Solution s;
  auto ret = s.longestPalindrome("babad");
  print_ret(ret);
  return 0;
}
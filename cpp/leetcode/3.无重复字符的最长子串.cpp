/*
 * @lc app=leetcode.cn id=3 lang=cpp
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (37.28%)
 * Likes:    5611
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.8M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 * 示例 4:
 *
 *
 * 输入: s = ""
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
#include <algorithm>
#include <iostream>
#include <string>
#include <unordered_map>
class Solution {
public:
  int lengthOfLongestSubstring(std::string s) {
    int ret = 0;
    int start = 0;
    std::unordered_map<char, int> pos;
    for (int i = 0; i < s.size(); i++) {
      char c = s.at(i);
      auto it = pos.find(c);
      if (it != pos.end() && it->second >= start) {
        start = it->second + 1;
      }
      pos[c] = i;
      ret = std::max(i - start + 1, ret);
    }
    return ret;
  }
};
// @lc code=end
void print_ret(int ret) {
  std::cout << "ret: ";
  std::cout << ret << " ";
  std::cout << std::endl;
}

int main(int argc, const char **argv) {
  Solution s;
  auto ret = s.lengthOfLongestSubstring("abcabcbb");
  print_ret(ret);
  return 0;
}

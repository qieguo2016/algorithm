/*
 * @lc app=leetcode.cn id=76 lang=cpp
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (45.17%)
 * Likes:    2496
 * Dislikes: 0
 * Total Accepted:    420.4K
 * Total Submissions: 930.7K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s
 * 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */

#include <iostream>
#include <string>
#include <unordered_map>
using namespace std;

// @lc code=start
class Solution {
public:
  // 双指针滑动窗口，右指针向前，覆盖之后左指针向前找最小值
  // 给t增加hashmap用作索引，另外再用一个hashmap存当前窗口各字符统计值
  string minWindow(string s, string t) {
    if (s.empty() || t.empty()) {
      return "";
    }
    int l = 0, r = -1, cnt = 0, m = s.size(), n = t.size();
    unordered_map<char, int> index, window;
    index.reserve(n);
    for (const auto ch : t) {
      index[ch]++;
    }
    int ret_l = -1, ret_len = m + 1; // 多个1，s全长刚好满足时可以满足>r-l+1
    // s = "ADOBECODEBANC", t = "ABC"
    while (r < m) {
      auto itir = index.find(s[++r]);
      if (itir != index.end()) {
        if (++window[s[r]] <= itir->second) { // 字符未冗余
          cnt++;
        }
      }
      while (cnt >= n && l <= r) { // 已覆盖，尝试精简
        if (ret_len > r - l + 1) { // 更新返回值
          ret_len = r - l + 1;
          ret_l = l;
        }

        auto itil = index.find(s[l]);
        if (itil != index.end()) {
          if (--window[s[l]] < itil->second) { // 已去掉冗余
            cnt--;
          }
        }
        l++;
      }
    }
    return ret_l == -1 ? "" : s.substr(ret_l, ret_len);
  }
};
// @lc code=end

int main(int argc, const char **argv) {
  Solution s;
  auto ret = s.minWindow("abc", "cba");
  cout << "ret:" << ret;
  return 0;
}

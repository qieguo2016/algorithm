/*
 * @lc app=leetcode.cn id=68 lang=cpp
 *
 * [68] 文本左右对齐
 *
 * https://leetcode.cn/problems/text-justification/description/
 *
 * algorithms
 * Hard (52.22%)
 * Likes:    334
 * Dislikes: 0
 * Total Accepted:    53.9K
 * Total Submissions: 103.2K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text",
 * "justification."]\n16'
 *
 * 给定一个单词数组 words
 * 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
 *
 * 你应该使用 “贪心算法”
 * 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 '
 * ' 填充，使得每行恰好有 maxWidth 个字符。
 *
 * 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
 *
 * 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
 *
 * 注意:
 *
 *
 * 单词是指由非空格字符组成的字符序列。
 * 每个单词的长度大于 0，小于等于 maxWidth。
 * 输入单词数组 words 至少包含一个单词。
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: words = ["This", "is", "an", "example", "of", "text",
 * "justification."], maxWidth = 16 输出:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 *
 * 示例 2:
 *
 *
 * 输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth =
 * 16 输出:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
 * 因为最后一行应为左对齐，而不是左右两端对齐。
 * ⁠ 第二行同样为左对齐，这是因为这行只包含一个单词。
 *
 *
 * 示例 3:
 *
 *
 * 输入:words =
 * ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth
 * = 20
 * 输出:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= words.length <= 300
 * 1 <= words[i].length <= 20
 * words[i] 由小写英文字母和符号组成
 * 1 <= maxWidth <= 100
 * words[i].length <= maxWidth
 *
 *
 */

#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  vector<string> fullJustify(vector<string> &words, int maxWidth) {
    vector<string> ret;
    if (words.empty()) {
      return ret;
    }
    int l = 0, r = 0, n = words.size();
    while (true) {
      int lineWidth = 0;
      while (r < n && lineWidth + r - l + words[r].length() <= maxWidth) {
        lineWidth += words[r++].length();
      }
      // last line
      if (r == n) {
        string s = join(words, l, r, " ");
        ret.emplace_back(s + blank(maxWidth - s.length()));
        break;
      }
      // one word
      if (r - l == 1) {
        ret.emplace_back(words[l] + blank(maxWidth - words[l].length()));
        l++;
        continue;
      }
      // more than one word
      int wordNum = r - l;
      int spaceNum = maxWidth - lineWidth;
      int avgSpace = spaceNum / (wordNum - 1); // segregation num:  wordNum - 1
      int extraSpace = spaceNum % (wordNum - 1);
      string s1 = join(words, l, l + extraSpace + 1, blank(avgSpace + 1));
      string s2 = join(words, l + extraSpace + 1, r, blank(avgSpace));
      ret.emplace_back(s1 + blank(avgSpace) + s2);
      l = r;
    }

    return ret;
  }

private:
  string blank(int n) { return string(n, ' '); }
  string join(vector<string> &words, int l, int r, string seg) {
    string ret(words[l]);
    for (size_t i = l + 1; i < r; i++) {
      ret += seg + words[i];
    }
    return ret;
  }
};
// @lc code=end

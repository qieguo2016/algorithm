/*
 * @lc app=leetcode.cn id=72 lang=cpp
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Hard (62.79%)
 * Likes:    2979
 * Dislikes: 0
 * Total Accepted:    369.4K
 * Total Submissions: 588.3K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数
 *  。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 和 word2 由小写英文字母组成
 *
 *
 */

#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // dp[i][j]表示s的前i个字符与t的前j个字符的编辑距离
  // 当s[i]==t[j]时，dp[i][j] = dp[i-1][j-1]
  // 当s[i]!=t[j]时，dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])+1
  // 其中dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1]
  // 表示插入操作
  int minDistance(string word1, string word2) {
    int m = word1.size(), n = word2.size();
    if (m * n == 0) {
      return m + n;
    }
    vector<vector<int>> dp(m + 1, vector<int>(n + 1));
    for (int i = 0; i <= m; i++) {
      dp[i][0] = i;
    }
    for (int i = 0; i <= n; i++) {
      dp[0][i] = i;
    }
    for (int i = 1; i <= m; i++) {
      for (int j = 1; j <= n; j++) {
        if (word1[i - 1] == word2[j - 1]) {
          dp[i][j] = dp[i - 1][j - 1];
        } else {
          dp[i][j] = min(dp[i - 1][j - 1], min(dp[i - 1][j], dp[i][j - 1])) + 1;
        }
      }
    }
    return dp[m][n];
  }
};
// @lc code=end

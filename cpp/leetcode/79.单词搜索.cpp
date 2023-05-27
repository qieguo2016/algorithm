/*
 * @lc app=leetcode.cn id=79 lang=cpp
 *
 * [79] 单词搜索
 *
 * https://leetcode.cn/problems/word-search/description/
 *
 * algorithms
 * Medium (46.29%)
 * Likes:    1601
 * Dislikes: 0
 * Total Accepted:    428.5K
 * Total Submissions: 925.7K
 * Testcase Example:
 * '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word
 * 存在于网格中，返回 true ；否则，返回 false 。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCCED"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "SEE"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCB"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n = board[i].length
 * 1
 * 1
 * board 和 word 仅由大小写英文字母组成
 *
 *
 *
 *
 * 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board
 * 更大的情况下可以更快解决问题？
 *
 */

#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // dfs+访问标记，标记矩阵可修改输入数组
  bool exist(vector<vector<char>> &board, string word) {
    if (board.empty()) {
      return word.empty();
    }

    for (int i = 0; i < board.size(); i++) {
      for (int j = 0; j < board[0].size(); j++) {
        if (dfs(board, i, j, word, 0)) {
          return true;
        }
      }
    }
    return false;
  }

private:
  bool dfs(vector<vector<char>> &board, int i, int j, string word, int idx) {
    if (idx >= word.size()) {
      return true;
    }
    if (i < 0 || i >= board.size() || j < 0 || j >= board[0].size()) {
      return false;
    }
    if (board[i][j] != word[idx]) {
      return false;
    }
    char ch = board[i][j];
    board[i][j] = '-';
    bool ret = dfs(board, i - 1, j, word, idx + 1) ||
               dfs(board, i + 1, j, word, idx + 1) ||
               dfs(board, i, j - 1, word, idx + 1) ||
               dfs(board, i, j + 1, word, idx + 1);
    board[i][j] = ch; // 回溯
    return ret;
  }
};
// @lc code=end

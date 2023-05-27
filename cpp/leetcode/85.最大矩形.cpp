/*
 * @lc app=leetcode.cn id=85 lang=cpp
 *
 * [85] 最大矩形
 *
 * https://leetcode.cn/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (54.65%)
 * Likes:    1528
 * Dislikes: 0
 * Total Accepted:    175.6K
 * Total Submissions: 321.4K
 * Testcase Example:
 * '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1
 * 的最大矩形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：6
 * 解释：最大矩形如上图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = []
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [["0"]]
 * 输出：0
 *
 *
 * 示例 4：
 *
 *
 * 输入：matrix = [["1"]]
 * 输出：1
 *
 *
 * 示例 5：
 *
 *
 * 输入：matrix = [["0","0"]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * rows == matrix.length
 * cols == matrix[0].length
 * 1 <= row, cols <= 200
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */
#include <iostream>
#include <stack>
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
  // 可以看成是84题的2维版本，0-n行的最大面积
  int maximalRectangle(vector<vector<char>> &matrix) {
    int m = matrix.size();
    if (m == 0) {
      return 0;
    }
    int n = matrix[0].size();
    vector<vector<int>> heights(m, vector<int>(n, 0));

    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (matrix[i][j] == '1') {
          heights[i][j] = (i == 0 ? 0 : heights[i - 1][j]) + 1;
        }
      }
    }

    int ret = 0;
    // 每一层按84题算一次
    for (int i = 0; i < m; i++) {
        ret = max(ret, largestRectangleArea(heights[i]));
    }
    return ret;
  }

private:
  // 84题的最大面积
  int largestRectangleArea(vector<int> &heights) {
    int n = heights.size();
    if (n < 2) {
      return n < 1 ? 0 : heights[0];
    }
    stack<int> stk; // 存位置
    stk.push(0);
    int ret = heights[0];
    for (int i = 1; i < heights.size(); i++) {
      while (!stk.empty() && heights[i] < heights[stk.top()]) {
        int h = heights[stk.top()];
        stk.pop();
        // 左侧是栈的下一个元素，右侧是i，栈空了说明左侧没有更小，也就左为0，宽度i-0，比如5/4/1/2的5和4
        int w = stk.empty() ? i : (i - stk.top() - 1);
        ret = max(ret, w * h);
      }
      stk.push(i);
    }
    while (!stk.empty()) {
      int h = heights[stk.top()];
      stk.pop();
      int w = stk.empty() ? n : (n - stk.top() - 1);
      ret = max(ret, w * h);
    }
    return ret;
  }
};
// @lc code=end

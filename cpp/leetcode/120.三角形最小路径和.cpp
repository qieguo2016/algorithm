/*
 * @lc app=leetcode.cn id=120 lang=cpp
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode.cn/problems/triangle/description/
 *
 * algorithms
 * Medium (68.67%)
 * Likes:    1214
 * Dislikes: 0
 * Total Accepted:    290.5K
 * Total Submissions: 423K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形 triangle ，找出自顶向下的最小路径和。
 *
 * 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与
 * 上一层结点下标 相同或者等于 上一层结点下标 + 1
 * 的两个结点。也就是说，如果正位于当前行的下标 i
 * ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * 输出：11
 * 解释：如下面简图所示：
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：triangle = [[-10]]
 * 输出：-10
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 *
 *
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
public:
  // 从底向上dp，dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + c[i][j]
  // 其中行[i+1]相同的，所以可以压缩到一维，每次循环直接叠加到列即可
  // 简化为dp[j] = min(dp[j], dp[j+1]) + c[i][j]
  int minimumTotal(vector<vector<int>> &triangle) {
    if (triangle.empty()) {
      return 0;
    }
    if (triangle.size() < 2) {
      return triangle[0][0];
    }
    vector<int> dp(
        triangle[triangle.size() - 1].begin(),
        triangle[triangle.size() - 1].end()); // 等腰三角形，列数等于行数
    for (int i = triangle.size() - 2; i >= 0; i--) {
      for (int j = 0; j < triangle[i].size(); j++) {
        dp[j] = min(dp[j], dp[j + 1]) + triangle[i][j];
      }
    }
    return dp[0];
  }
};
// @lc code=end

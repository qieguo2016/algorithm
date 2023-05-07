/*
 * @lc app=leetcode.cn id=407 lang=cpp
 *
 * [407] 接雨水 II
 *
 * https://leetcode.cn/problems/trapping-rain-water-ii/description/
 *
 * algorithms
 * Hard (57.80%)
 * Likes:    664
 * Dislikes: 0
 * Total Accepted:    33.9K
 * Total Submissions: 58.6K
 * Testcase Example:  '[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]'
 *
 * 给你一个 m x
 * n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
 * 输出: 4
 * 解释: 下雨后，雨水将会被上图蓝色的方块中。总的接雨水量为1+2+1=4。
 *
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入: heightMap =
 * [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
 * 输出: 10
 *
 *
 *
 *
 * 提示:
 *
 *
 * m == heightMap.length
 * n == heightMap[i].length
 * 1 <= m, n <= 200
 * 0 <= heightMap[i][j] <= 2 * 10^4
 *
 *
 *
 *
 */

#include <algorithm>
#include <array>
#include <functional>
#include <queue>
#include <utility>
#include <vector>
// @lc code=start
class Solution {
  using pii = std::pair<int, int>;

public:
  int trapRainWater(std::vector<std::vector<int>> &heightMap) {
    int res = 0;
    if (heightMap.size() < 3 || heightMap[0].size() < 3) {
      return res;
    }
    int m = heightMap.size();
    int n = heightMap[0].size();
    std::vector<bool> visited(m * n, false);
    std::priority_queue<pii, std::vector<pii>, std::greater<pii>> q;

    // check edge
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (i == 0 || i == m - 1 || j == 0 || j == n - 1) {
          int idx = i * n + j;
          q.push({heightMap[i][j], idx});
          visited[idx] = true;
        }
      }
    }

    int dirs[] = {-1, 0, 1, 0, -1};
    // check every grid
    while (!q.empty()) {
      auto cur = q.top();
      q.pop();

      for (int k = 0; k < 4; k++) { // check left/down/right/up
        int x = cur.second / n + dirs[k];
        int y = cur.second % n + dirs[k + 1];
        int idx = x * n + y;
        if (x >= 0 && x < m && y >= 0 && y < n && !visited[idx]) {
          res += cur.first > heightMap[x][y] ? cur.first - heightMap[x][y] : 0;
          visited[idx] = true;
          q.push({std::max(heightMap[x][y], cur.first), idx});
        }
      }
    }
    return res;
  }
};
// @lc code=end
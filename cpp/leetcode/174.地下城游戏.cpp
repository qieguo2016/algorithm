/*
 * @lc app=leetcode.cn id=174 lang=cpp
 *
 * [174] 地下城游戏
 *
 * https://leetcode.cn/problems/dungeon-game/description/
 *
 * algorithms
 * Hard (48.76%)
 * Likes:    728
 * Dislikes: 0
 * Total Accepted:    62.6K
 * Total Submissions: 128.4K
 * Testcase Example:  '[[-2,-3,3],[-5,-10,1],[10,30,-5]]'
 *
 * table.dungeon, .dungeon th, .dungeon td {
 * ⁠ border:3px solid black;
 * }
 *
 * ⁠.dungeon th, .dungeon td {
 * ⁠   text-align: center;
 * ⁠   height: 70px;
 * ⁠   width: 70px;
 * }
 *
 * 恶魔们抓住了公主并将她关在了地下城 dungeon 的 右下角 。地下城是由 m x n
 * 个房间组成的二维网格。我们英勇的骑士最初被安置在 左上角
 * 的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
 *
 * 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0
 * 或以下，他会立即死亡。
 *
 * 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为
 * 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
 *
 * 为了尽快解救公主，骑士决定每次只 向右 或 向下 移动一步。
 *
 * 返回确保骑士能够拯救到公主所需的最低初始健康点数。
 *
 * 注意：任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
 * 输出：7
 * 解释：如果骑士遵循最佳路径：右 -> 右 -> 下 -> 下 ，则骑士的初始健康点数至少为
 * 7 。
 *
 * 示例 2：
 *
 *
 * 输入：dungeon = [[0]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == dungeon.length
 * n == dungeon[i].length
 * 1 <= m, n <= 200
 * -1000 <= dungeon[i][j] <= 1000
 *
 *
 */

#include <limits.h>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  int calculateMinimumHP(vector<vector<int>> &dungeon) {
    // 粗看是dp题，从终点往回思考归纳出状态转移方程，输出一般有两种结果，1是终点/起点值，2是维护中间变量
    // 先看状态转移方程，倒数第二步可以是[n-1][n]或者[n][n-1]两种，应该选初始值更小的路径
    // 设dp[i][j]是从左上角走到i,j的最小初始值，
    // dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), sum(x,y))
    // dp还需要看走左上或者右上的sub来判断，没法做dp，可以考虑从后往前dp去掉sum这一项

    // 设dp[i][j]是从i,j到终点所需最小初始值,
    // 对于终点，如果是正数取1，负数则取1-c[m][n]，也就是max(1-c[m][n], 1)
    // 对于边界[m-1][n]，dp=max(dp[m][n]-c[m][n], 1)，(初始值最小也是1)
    // 其他dp[i][j] = max(min(dp[i+1][j], dp[i][j+1]) - c[i][j]), 1)
    // 为了将矩阵边界也统一到上式，可以将边界设置为max，那min的时候就会取到下或者右

    int m = dungeon.size(), n = dungeon[0].size();
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, INT_MAX));
    dp[m - 1][n] = dp[m][n - 1] = 1; // dp[m-1][n-1] = max(1-dungeon[i][j], 1)
    for (int i = m - 1; i >= 0; i--) {
      for (int j = n - 1; j >= 0; j--) {
        dp[i][j] = max(min(dp[i + 1][j], dp[i][j + 1]) - dungeon[i][j], 1);
      }
    }
    return dp[0][0];
  }
};
// @lc code=end

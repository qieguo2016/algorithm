/*
 * @lc app=leetcode.cn id=56 lang=cpp
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (45.79%)
 * Likes:    967
 * Dislikes: 0
 * Total Accepted:    245K
 * Total Submissions: 535K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti,
 * endi]
 * 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * intervals[i].length == 2
 * 0 i i
 *
 *
 */

#include <vector>

// @lc code=start
class Solution {
public:
  std::vector<std::vector<int>>
  merge(std::vector<std::vector<int>> &intervals) {
    std::vector<std::vector<int>> ret;
    if (intervals.empty()) {
      return ret;
    }
    sort(intervals.begin(), intervals.end());
    std::vector<int> pre = intervals[0];
    for (size_t i = 1; i < intervals.size(); i++) {
      // [[1,3],[2,6],[8,10],[15,18]]
      if (intervals[i][0] <= pre[1]) {
        pre[1] = std::max(pre[1], intervals[i][1]);
      } else {
        ret.emplace_back(pre);
        pre = intervals[i];
      }
    }
    ret.emplace_back(pre);
    return ret;
  }
};
// @lc code=end

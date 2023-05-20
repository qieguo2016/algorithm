/*
 * @lc app=leetcode.cn id=210 lang=cpp
 *
 * [210] 课程表 II
 *
 * https://leetcode.cn/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (56.59%)
 * Likes:    774
 * Dislikes: 0
 * Total Accepted:    177.2K
 * Total Submissions: 313K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses -
 * 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi]
 * ，表示在选修课程 ai 前 必须 先选修 bi 。
 *
 *
 * 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
 *
 *
 * 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回
 * 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：[0,1]
 * 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程
 * 0。因此，正确的课程顺序为 [0,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * 输出：[0,2,1,3]
 * 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1
 * 和课程 2 都应该排在课程 0 之后。 因此，一个正确的课程顺序是 [0,1,2,3]
 * 。另一个正确的排序是 [0,2,1,3] 。
 *
 * 示例 3：
 *
 *
 * 输入：numCourses = 1, prerequisites = []
 * 输出：[0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * ai != bi
 * 所有[ai, bi] 互不相同
 *
 *
 */

#include <deque>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 按dag执行一遍
  vector<int> findOrder(int numCourses, vector<vector<int>> &prerequisites) {
    if (numCourses <= 0) {
      return vector<int>();
    }
    if (numCourses <= 1) {
      return vector<int>(numCourses, 0);
    }
    vector<vector<int>> graph(numCourses, vector<int>(0)); // output
    vector<int> indegree(numCourses, 0);
    for (int i = 0; i < prerequisites.size(); i++) {
      graph[prerequisites[i][1]].push_back(prerequisites[i][0]);
      indegree[prerequisites[i][0]]++;
    }

    deque<int> q;
    for (int i = 0; i < numCourses; i++) {
      if (indegree[i] < 1) {
        q.push_back(i);
      }
    }

    vector<int> order;
    order.reserve(numCourses);
    while (!q.empty()) {
      int idx = q.front();
      q.pop_front();
      order.push_back(idx);
      for (int i = 0; i < graph[idx].size(); i++) {
        int j = graph[idx][i];
        indegree[j]--;
        if (indegree[j] <= 0) {
          q.push_back(j);
        };
      }
    }
    return numCourses == order.size() ? order : vector<int>();
  }
};
// @lc code=end

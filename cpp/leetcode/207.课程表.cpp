/*
 * @lc app=leetcode.cn id=207 lang=cpp
 *
 * [207] 课程表
 *
 * https://leetcode.cn/problems/course-schedule/description/
 *
 * algorithms
 * Medium (53.58%)
 * Likes:    1595
 * Dislikes: 0
 * Total Accepted:    297.9K
 * Total Submissions: 556K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites
 * 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须
 * 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1
 * 之前，你需要先完成​课程 0 ；并且学习课程 0
 * 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * prerequisites[i].length == 2
 * 0 i, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */

#include <deque>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  // 检查图是否有环，可按dag执行一遍，如果未全执行完则有环
  bool canFinish(int numCourses, vector<vector<int>> &prerequisites) {
    if (numCourses <= 1) {
      return true;
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

    int cnt = 0;
    while (!q.empty()) {
      int idx = q.front();
      q.pop_front();
      cnt++;
      for (int i = 0; i < graph[idx].size(); i++) {
        int j = graph[idx][i];
        indegree[j]--;
        if (indegree[j] <= 0) {
          q.push_back(j);
        };
      }
    }
    return cnt == numCourses;
  }
};
// @lc code=end

/*
 * @lc app=leetcode.cn id=84 lang=cpp
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (44.98%)
 * Likes:    2446
 * Dislikes: 0
 * Total Accepted:    342.6K
 * Total Submissions: 761.7K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为
 * 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：heights = [2,1,5,6,2,3]
 * 输出：10
 * 解释：最大的矩形为图中红色区域，面积为 10
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入： heights = [2,4]
 * 输出： 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
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
  // 每个高度对应的最大面积由宽度决定，宽度则为左右第一个比该高度小所包起来的区间
  // 举例[8,2,4,5,3,2], 直接看245，到3时则可以确定5的最大面积是5，然后4是4*2
  // 2则还不能确定, 所以，发现递增的时候先记录起来，当发现比最新值小，则可以计算
  // 显然可以用栈来记录，也就是单调栈
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

int main(int argc, const char **argv) {
  Solution s;
  vector<int> arr({5, 4, 1, 2});
  auto ret = s.largestRectangleArea(arr);
  std::cout << "ret: " << ret;
  return 0;
}
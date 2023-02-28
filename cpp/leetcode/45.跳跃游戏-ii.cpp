/*
 * @lc app=leetcode.cn id=45 lang=cpp
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Medium (40.39%)
 * Likes:    1013
 * Dislikes: 0
 * Total Accepted:    139.1K
 * Total Submissions: 344.3K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 *
 * 假设你总是可以到达数组的最后一个位置。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,3,0,1,4]
 * 输出: 2
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * 0
 *
 *
 */

#include <algorithm>
#include <vector>
// 当前和下一步合并起来最大即为最优，证明：
//   第1步可以走到[1,n],第2步分别可以走的最大距离是x1,x2..,根据第2步的最大距离可以选出第1步的选择
//   是否存在第1步不走上面的选择但是整体最优呢？不存在，因为第2步的最大距离是覆盖了非最优解的，步数会更小。
// @lc code=start
class Solution {
public:
  int jump(std::vector<int> &nums) {
    if (nums.size() <= 1) {
      return 0; // 不用走就到末尾了
    }
    int preMax = nums[0]; // 上一步的最大位置
    int curMax = preMax;  // 当前步的最大位置
    int step = 1;         // 先走一步
    int idx = 1;
    while (preMax < nums.size() - 1) { // 最后不用走，算进来会多走一步
      // [4,1,3,2,1,2,8]
      curMax = std::max(curMax, idx + nums[idx]);
      if (idx >= preMax) { // 到了上一步的最大距离，可以确定上一步的选择
        preMax = curMax;
        step++;
      }
      idx++;
    }
    return step;
  }
};
// @lc code=end

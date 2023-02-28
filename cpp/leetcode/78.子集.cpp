/*
 * @lc app=leetcode.cn id=78 lang=cpp
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (79.88%)
 * Likes:    1214
 * Dislikes: 0
 * Total Accepted:    258.2K
 * Total Submissions: 323.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同
 * 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */

#include <vector>

// @lc code=start
class Solution {
private:
  std::vector<std::vector<int>> subsetsAppend(std::vector<int> &nums) {
    std::vector<std::vector<int>> ret;
    ret.push_back({});
    for (auto num : nums) {
      int size = ret.size();
      for (int i = 0; i < size; i++) {
        std::vector<int> arr = ret[i];
        arr.push_back(num);
        ret.push_back(arr);
      }
    }
    return ret;
  }

public:
  std::vector<std::vector<int>> subsets(std::vector<int> &nums) {
    return subsetsAppend(nums);
  }
};
// @lc code=end

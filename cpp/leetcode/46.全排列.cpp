/*
 * @lc app=leetcode.cn id=46 lang=cpp
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.89%)
 * Likes:    2396
 * Dislikes: 0
 * Total Accepted:    796K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序
 * 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */

#include <vector>

// @lc code=start
class Solution {
private:
  void backtrack(std::vector<std::vector<int>> &res, std::vector<int> &nums,
                 int cur) {
    if (cur == nums.size()) {
      res.emplace_back(nums);
      return;
    }
    // 
    for (int i = cur; i < nums.size(); i++) {
      std::swap(nums[i], nums[cur]);
      backtrack(res, nums, cur + 1);
      std::swap(nums[cur], nums[i]);
    }
  }

  void backtrack(std::vector<std::vector<int>> &res,
                 const std::vector<int> &nums, std::vector<int> &out,
                 std::vector<bool> &visited) {
    if (out.size() == nums.size()) {
      res.emplace_back(out);
      return;
    }

    for (int i = 0; i < nums.size(); i++) {
      if (visited[i]) {
        continue;
      }
      visited[i] = true;
      out.emplace_back(nums[i]);
      backtrack(res, nums, out, visited);
      out.pop_back();
      visited[i] = false;
    }
  }

public:
  std::vector<std::vector<int>> permute(std::vector<int> &nums) {
    std::vector<std::vector<int>> res;
    backtrack(res, nums, 0);
    return res;

    // std::vector<std::vector<int>> res;
    // std::vector<bool> visited(nums.size());
    // std::vector<int> out;
    // backtrack(res, nums, out, visited);
    // return res;
  }
};
// @lc code=end

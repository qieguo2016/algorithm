/*
 * @lc app=leetcode.cn id=39 lang=cpp
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.46%)
 * Likes:    1383
 * Dislikes: 0
 * Total Accepted:    269.4K
 * Total Submissions: 371.8K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 *
 *
 * 示例 1：
 *
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2：
 *
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 *
 *
 */

#include <algorithm>
#include <vector>

// @lc code=start
class Solution {
private:
  void backtrack(const std::vector<int> &candidates, int target, int start,
                 std::vector<int> &cur, std::vector<std::vector<int>> &out) {
    if (target < 0) {
      return;
    }
    if (target == 0) {
      out.push_back(cur);
      return;
    }

    for (size_t i = start; i < candidates.size(); i++) {
      cur.push_back(candidates[i]);
      backtrack(candidates, target - candidates[i], i, cur, out);
      cur.pop_back();
    }
  }

public:
  std::vector<std::vector<int>> combinationSum(std::vector<int> &candidates,
                                               int target) {
    std::sort(candidates.begin(), candidates.end());
    std::vector<int> cur;
    std::vector<std::vector<int>> out;
    backtrack(candidates, target, 0, cur, out);
    return out;
  }
};
// @lc code=end

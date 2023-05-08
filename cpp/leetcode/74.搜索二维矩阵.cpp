/*
 * @lc app=leetcode.cn id=74 lang=cpp
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode.cn/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (48.61%)
 * Likes:    795
 * Dislikes: 0
 * Total Accepted:    308.4K
 * Total Submissions: 634.3K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x
 * n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -10^4
 *
 *
 */

#include <algorithm>
#include <vector>

// @lc code=start
class Solution {
public:
  bool searchMatrix(std::vector<std::vector<int>> &matrix, int target) {
    // 第一个大于目标值的列
    auto t_row = std::upper_bound(
        matrix.begin(), matrix.end(), target,
        [](const int t, const std::vector<int> &row) { return t < row[0]; });
    if (t_row == matrix.begin()) {
      return false;
    }
    --t_row;
    return std::binary_search(t_row->begin(), t_row->end(), target);
  }
};
// @lc code=end

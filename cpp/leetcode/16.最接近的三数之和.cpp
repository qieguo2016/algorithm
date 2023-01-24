/*
 * @lc app=leetcode.cn id=16 lang=cpp
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (40.62%)
 * Likes:    187
 * Dislikes: 0
 * Total Accepted:    26.7K
 * Total Submissions: 65.8K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和
 * 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */

// @lc code=start
#include <algorithm>
#include <iostream>
#include <vector>

class Solution {
public:
  int threeSumClosest(std::vector<int> &nums, int target) {
    int minDiff = INT_MAX;
    int closest = 0;
    sort(nums.begin(), nums.end());
    for (int i = 0; i < nums.size(); i++) {
      int j = i + 1, k = nums.size() - 1;
      while (j < k) {
        int threeSum = nums[i] + nums[j] + nums[k];
        int diff = abs(target - threeSum);
        if (diff == 0) {
          return 0;
        }
        if (diff < minDiff) {
          minDiff = diff;
          closest = threeSum;
        }
        if (threeSum > target) {
          k--;
        } else {
          j++;
        }
      }
    }
    return closest;
  }
};
// @lc code=end

int main(int argc, const char **argv) {
  Solution s;
  std::vector<int> nums = {2,3,4,5};
  auto str = s.threeSumClosest(nums, 17);
  std::cout << str << std::endl;
  return 0;
}
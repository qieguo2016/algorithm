/*
 * @lc app=leetcode.cn id=15 lang=cpp
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (32.32%)
 * Likes:    3410
 * Dislikes: 0
 * Total Accepted:    532.8K
 * Total Submissions: 1.6M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c
 * ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */

// @lc code=start
#include <algorithm>
#include <iostream>
#include <ostream>
#include <vector>
class Solution {
public:
  std::vector<std::vector<int>> threeSum(std::vector<int> &nums) {
    std::vector<std::vector<int>> ret;
    if (nums.size() < 3) {
      return ret;
    }
    std::sort(nums.begin(), nums.end());
    if (nums.at(nums.size() - 1) < 0) {
      return ret;
    }
    for (size_t i = 0; i < nums.size() - 2; i++) {
      if (nums.at(0) > 0) {
        break;
      }
      if (i > 0 && nums.at(i) == nums.at(i - 1)) {
        continue;
      }
      int rest = 0 - nums.at(i);
      int l = i + 1, r = nums.size() - 1;
      while (l < r) {
        int tmp = nums.at(l) + nums.at(r);
        if (tmp == rest) {
          ret.push_back({nums.at(i), nums.at(l), nums.at(r)});
          while (l < r && nums.at(l) == nums.at(l + 1)) {
            l++;
          }
          while (l < r && nums.at(r) == nums.at(r - 1)) {
            r--;
          }
          l++;
          r--;
        } else if (tmp < rest) {
          l++;
        } else {
          r--;
        }
      }
    }
    return ret;
  }
};
// @lc code=end

void print_ret(const std::vector<std::vector<int>> &ret) {
  std::cout << "ret: [" << std::endl;
  for (size_t i = 0; i < ret.size(); i++) {
    std::cout << "[";
    for (size_t j = 0; j < ret.at(i).size(); j++) {
      std::cout << ret.at(i).at(j) << ",";
    }
    std::cout << "]" << std::endl;
  }
  std::cout << "]" << std::endl;
}

int main(int argc, const char **argv) {
  Solution s;
  std::vector<int> arr = {-1, 0, 1, 2, -1, -4};
  auto ret = s.threeSum(arr);
  print_ret(ret);
  return 0;
}
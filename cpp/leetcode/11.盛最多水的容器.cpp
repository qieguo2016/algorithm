/*
 * @lc app=leetcode.cn id=11 lang=cpp
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (63.82%)
 * Likes:    2531
 * Dislikes: 0
 * Total Accepted:    454.7K
 * Total Submissions: 712.4K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai)
 * 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)
 * 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 解释：图中垂直线代表输入数组
 * [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 * 示例 2：
 *
 *
 * 输入：height = [1,1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：height = [4,3,2,1,4]
 * 输出：16
 *
 *
 * 示例 4：
 *
 *
 * 输入：height = [1,2,1]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * n = height.length
 * 2
 * 0
 *
 *
 */

// @lc code=start
#include <algorithm>
#include <iostream>
#include <vector>
class Solution {
public:
  int maxArea(std::vector<int> &height) {
    int l = 0, r = height.size() - 1, ma = 0;
    while (l < r) {
      if (height.at(l) < height.at(r)) {
        ma = std::max(ma, (r - l) * height.at(l));
        l++;
      } else {
        ma = std::max(ma, (r - l) * height.at(r));
        r--;
      }
    }
    return ma;
  }
};
// @lc code=end

void print_ret(int ret) {
  std::cout << "ret: ";
  std::cout << ret << " ";
  std::cout << std::endl;
}

int main(int argc, const char **argv) {
  Solution s;
  std::vector<int> arr = {1,8,6,2,5,4,8,3,7};
  auto ret = s.maxArea(arr);
  print_ret(ret);
  return 0;
}
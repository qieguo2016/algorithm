/*
 * @lc app=leetcode.cn id=128 lang=cpp
 *
 * [128] 最长连续序列
 *
 * https://leetcode.cn/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (54.92%)
 * Likes:    1645
 * Dislikes: 0
 * Total Accepted:    385.8K
 * Total Submissions: 703.1K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组 nums
 * ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 *
 * 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [100,4,200,1,3,2]
 * 输出：4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^9
 *
 *
 */
#include <set>
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
  // 朴素解法：遍历数组，对每一个数不断检查递增值，更新最大长度
  // 优化：检查递增时可以用hash set保存是否存在
  // 优化2：检查递增值时会有重复检查，可设法只让最小值判断一次。
  //       可判断-1值是否存在，如果存在则让-1去检查，直到最小值才真正循环起来判断
  int longestConsecutive(vector<int> &nums) {
    if (nums.empty()) {
      return 0;
    }
    int ret = 1;
    set<int> exist(nums.begin(), nums.end());
    for (auto num : nums) {
      if (exist.find(num - 1) != exist.end()) {
        continue;
      }
      int cnt = 1;
      int cur = num + 1;
      while (exist.find(cur) != exist.end()) {
        cnt++;
        cur++;
      }
      ret = max(ret, cnt);
    }
    return ret;
  }
};
// @lc code=end

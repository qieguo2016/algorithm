/*
 * @lc app=leetcode.cn id=300 lang=cpp
 *
 * [300] 最长递增子序列
 *
 * https://leetcode.cn/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (54.72%)
 * Likes:    3198
 * Dislikes: 0
 * Total Accepted:    732.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 *
 * 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7]
 * 是数组 [0,3,1,6,2,2,7] 的子序列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,9,2,5,3,7,101,18]
 * 输出：4
 * 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,3,2,3]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,7,7,7,7,7,7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
 *
 *
 */

#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  int lengthOfLIS(vector<int> &nums) { return nLogN(nums); }

private:
  // dp题目，过程中维护最长子序列，人工分析规律来找状态转移方程
  // [9,2,5,3,10,12,6,7,13]，
  // 遍历到3时，可以跟2组成23，遍历到10时，有2310和2510两种可能，12则可以跟10连起来
  // 遍历到6，有256和236两种，7也可以跟6的连起来，遍历到13的时候，可以跟12也可以跟7连起来
  // 因此，可以看到每个元素都要跟之前的每个元素比较一下，从而得到包含该元素的最长序列
  // 从而可定义dp[i]是包含该元素的最长子序列，dp[i]=max(dp[j]...)+1,
  // 其中0<=j<i且nums[j]<nums[i] 所以结果就是max(dp[i]...),
  // 遍历每个元素时往回遍历所有元素，复杂度是O(n^2)
  int n2(vector<int> &nums) {
    int n = nums.size();
    if (n <= 1) {
      return n;
    }
    int ret = 1;
    vector<int> dp(n, 1);
    for (int i = 1; i < n; i++) {
      for (int j = i - 1; j >= 0; j--) {
        dp[i] = nums[i] > nums[j] ? max(dp[i], dp[j] + 1) : dp[i];
      }
      ret = max(ret, dp[i]);
    }
    return ret;
  }

  // 上面的例子里面5和3其实可以直接丢弃5了，因为同等长度下，3更小，未来有更多可能，最小末尾数是3。
  // 同理，遍历到7时，长度为4的子序列其实就是2367了，最小末尾数是7，10和12都可以丢弃了，后面不会再用到。
  // 而遍历每个元素的时候，当前元素如果比长度为1的最小末尾数大，那可以组成长度为2的子序列，
  // 再依次比较234...n的子序列，n为当前最长递增长度，就可以知道当前可以组成的最长子序列
  int nLogN(vector<int> &nums) {
    int n = nums.size();
    if (n <= 1) {
      return n;
    }
    int ret = 1;
    vector<int> tails(n, 0);
    tails[0] = nums[0];
    for (int i = 1; i < nums.size(); i++) {
      int num = nums[i];
      if (num > tails[ret - 1]) {
        tails[ret++] = num;
        continue;
      }
      int l = 0, r = ret;
      while (l < r) {
        int mid = l + (r - l) / 2;
        if (num > tails[mid]) {
          l = mid + 1;
        } else {
          r = mid;
        }
      }
      tails[l] = num;
    }
    return ret;
  }
};
// @lc code=end

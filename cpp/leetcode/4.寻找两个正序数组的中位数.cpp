/*
 * @lc app=leetcode.cn id=4 lang=cpp
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (41.62%)
 * Likes:    6156
 * Dislikes: 0
 * Total Accepted:    863.9K
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1
 * 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
#include <algorithm>
#include <iostream>
#include <vector>
class Solution {
public:
  double findMedianSortedArrays(std::vector<int> &nums1,
                                std::vector<int> &nums2) {
    int total = nums1.size() + nums2.size(); // total >= 1
    if ((total) % 2 == 1) {
      return (double)findKth(nums1, nums2, total / 2);
    }
    int left = total / 2 - 1;
    int right = total / 2;
    return (findKth(nums1, nums2, left) + findKth(nums1, nums2, right)) / 2.0;
  }

private:
  int findKth(std::vector<int> &nums1, std::vector<int> &nums2, int k) {
    int l1 = nums1.size();
    int l2 = nums2.size();
    int i = 0, j = 0;

    while (true) {
      if (i >= l1) {
        return nums2[j + k];
      }
      if (j >= l2) {
        return nums1[i + k];
      }
      if (k == 0) {
        return std::min(nums1[i], nums2[j]);
      }

      // 如果k/2超过了某一个队列的长度，由于k是合法输入，那第k个不能在长队列的小端
      // 反推一下，假如k在长队列的小端，那么也要从短队列中搬运超过k/2的元素过来，但是k/2大于队列长度，不成立
      // k/2都在数组中的情况，这时肯定不可能在最小端，推理同上
      int ii = std::min(i + (k - 1) / 2, l1 - 1);  
      int jj = std::min(j + (k - 1) / 2, l2 - 1);
      // 数组序号是从0开始的，取中位数时要用(k-1)/2，否则偶数个数会有1个位置右偏，而丢弃时再+1，并不是居中丢弃
      // 比如4/2=2是第三个数，再+1丢弃了3个，按理二分应该丢2个
      if (nums1[ii] <= nums2[jj]) {
        k -= ii - i + 1; // 扔掉了ii - i + 1个
        i = ii + 1;
      } else {
        k -= jj - j + 1;
        j = jj + 1;
      }
    }
  }
};
// @lc code=end

void print_ret(double ret) {
  std::cout << "ret: ";
  std::cout << ret << " ";
  std::cout << std::endl;
}

int main(int argc, const char **argv) {
  Solution s;
  std::vector<int> a1 = {1,2};
  std::vector<int> a2 = {3,4};
  auto ret = s.findMedianSortedArrays(a1, a2);
  print_ret(ret);
  return 0;
}

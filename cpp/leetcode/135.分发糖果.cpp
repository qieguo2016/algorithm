/*
 * @lc app=leetcode.cn id=135 lang=cpp
 *
 * [135] 分发糖果
 *
 * https://leetcode.cn/problems/candy/description/
 *
 * algorithms
 * Hard (50.54%)
 * Likes:    1196
 * Dislikes: 0
 * Total Accepted:    221K
 * Total Submissions: 437K
 * Testcase Example:  '[1,0,2]'
 *
 * n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
 *
 * 你需要按照以下要求，给这些孩子分发糖果：
 *
 *
 * 每个孩子至少分配到 1 个糖果。
 * 相邻两个孩子评分更高的孩子会获得更多的糖果。
 *
 *
 * 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：ratings = [1,0,2]
 * 输出：5
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
 *
 *
 * 示例 2：
 *
 *
 * 输入：ratings = [1,2,2]
 * 输出：4
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
 * ⁠    第三个孩子只得到 1
 * 颗糖果，这满足题面中的两个条件。
 *
 *
 *
 * 提示：
 *
 *
 * n == ratings.length
 * 1 <= n <= 2 * 10^4
 * 0 <= ratings[i] <= 2 * 10^4
 *
 *
 */
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
  // 朴素解法，两次遍历，准备两个糖果数组。一次发现递增区间则对应位置糖果递增，否则置为1；
  // 再一次从右向左遍历，发现递增少区间也是对应位置糖果加1，这次其实是看递减的，取两个数组同位置较大者即可。
  // 优化：一次遍历也可以发现递减趋势，而且递减区间和递增是可以对应起来的，另外题目要求取总数，可以标量存总数
  //    举例 1,2,3,3,2,1,2，对应糖果是 1,2,4,3,2,1,2。 3号是递减开始所以是4
  //    观察321，等价于123，为了保证降序，还需要把前一个3+1改成4，也就是再加（递减个数-递增个数）
  int candy(vector<int> &ratings) {
    if (ratings.empty()) {
      return 0;
    }
    int ret = 1;                   // 第一个先按给1个糖果算
    int inc = 1, dec = 0, pre = 1; // 递增递减区间个数，pre是前一个同学的糖果数
    for (int i = 1; i < ratings.size(); i++) {
      if (ratings[i] >= ratings[i - 1]) {
        pre = ratings[i] == ratings[i - 1]
                  ? 1
                  : 1 + pre; // 相等也是不大于，可以只分1个
        ret += pre;
        dec = 0; // 重置递减区间
        inc = pre;
      } else {
        dec++;
        ret += dec;
        if (dec >= inc) { // 再加（递减个数-递增个数），每次循环加1个
          ret++;
        }
        pre = 1; // 重置递减区间
      }
    }
    return ret;
  }
};
// @lc code=end

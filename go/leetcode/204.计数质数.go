/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (38.19%)
 * Likes:    702
 * Dislikes: 0
 * Total Accepted:    150.5K
 * Total Submissions: 394.1K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 10
 * 输出：4
 * 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 * 示例 2：
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 * 示例 3：
 *
 * 输入：n = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 5 * 10^6
 *
 *
 */

// @lc code=start
func countPrimes(n int) int {
	// 1. 对于每一个数，从2开始到根号n试探是否能被整除，到根号n是因为大于根号n之后，另外一个因子也必然小于根号n
	// 2. 假如一个数是质数，那2x,3x就不再是质数，可以递推一直标记到n
	// 3. 2中的标记会有重复，比如24，2、3、4、6、8、12都会算一次，可以用线性筛优化，具体不展开
	composite := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !composite[i] {
			count++
			for j := i + i; j < n; j += i {
				composite[j] = true
			}
		}
	}
	return count
}

// @lc code=end


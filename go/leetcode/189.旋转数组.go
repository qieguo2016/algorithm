/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Medium (45.57%)
 * Likes:    983
 * Dislikes: 0
 * Total Accepted:    270.4K
 * Total Submissions: 593.3K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 *
 *
 * 进阶：
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,3,4,5,6,7], k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 *
 * 输入：nums = [-1,-100,3,99], k = 2
 * 输出：[3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 * 0
 *
 *
 *
 *
 *
 */

// @lc code=start
func rotate(nums []int, k int) {
	// 最直观的想法，类似ringbuf，超出位置对len取余放到头部即可，时间O(n)、空间O(n)
	// ret := make([]int, len(nums))
	// for i := 0; i < len(nums); i++ {
	// 	ret[(i+k)%len(nums)] = nums[i]
	// }
	// copy(nums, ret)

	// 上面的方式只要一个变量存储要移位的值，每次只处理一个位置
	// 比如 [1,2,3,4,5,6]到[5,6,1,2,3,4]，1到3，3到5，5到1
	// 一直走的话，会形成循环，第一次回到原点的走过的路程就是nk的最小公倍数
	// 那么单次处理的元素个数=最小公倍数/k，要遍历所有元素，那么循环次数=nk/最小公倍数，也就是最大公约数
	// 时间复杂度每个元素最多遍历一次，所以是O(n), 空间O(1)
	//
	// n := len(nums)
	// k %= n
	// for i := 0; i < gcd(k, n); i++ { // 循环次数
	// 	pre, cur := nums[i], i
	// 	for { // 每组轮换到原点
	// 		next := (cur + k) % n             // 下一跳
	// 		nums[next], pre = pre, nums[next] // 交换13
	// 		cur = next
	// 		if cur == i { // 单次循环到原点，跳出
	// 			break
	// 		}
	// 	}
	// }

	// 反转法比较难想到，[1,2,3,4,5,6,7] 到 [5,6,7,1,2,3,4]
	// 可以看成是 [7,6,5,4,3,2,1]的分组反转 先反转 567， 再反转1234
	// 遍历两次数组，O(n)，空间O(1)
	n := len(nums)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	k %= n
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func gcd(a, b int) int { // 最大公约数
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end


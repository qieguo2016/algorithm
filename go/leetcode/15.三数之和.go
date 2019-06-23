/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (22.51%)
 * Likes:    1015
 * Dislikes: 0
 * Total Accepted:    58.1K
 * Total Submissions: 258.1K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */

import (
	"sort"
)

// 1. 先排序
// 2. 固定一个，后两个数从边缘往中间搜索
// 3. 做一些减枝优化
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	// 三元组，所以只需要遍历到倒数第三个
	for i := 0; i < len(nums) - 2; i++ {
		// 剪枝优化1，最小值>0则退出
		if nums[i] > 0 {
			break
		}
		// 去重1，与前面相同则不再搜索
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 缓存target
		target := 0 - nums[i]
		// 双指针分别指向两侧边缘
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[j] + nums[k] == target {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				// 去重2
				for j < k && nums[j] == nums[j + 1] {
					j++
				}
				// 去重3
				for j < k && nums[k] == nums[k - 1] {
					k--
				}
				j++
				k--
			} else if nums[j] + nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}


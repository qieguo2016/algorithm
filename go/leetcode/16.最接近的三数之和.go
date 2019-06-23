/*
 * @lc app=leetcode.cn id=16 lang=golang
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
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 * 
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 * 
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 * 
 * 
 */

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	// 默认有解：数组长度为3以上
	minDiff := math.MaxFloat64
	closest := 0
	sort.Ints(nums)
	println(nums)
	l := len(nums)
	for i := 0; i < l - 2; i++ {
		j, k := i+1, l-1
		for j < k {
			threeSum := nums[i] + nums[j] + nums[k]
			diff := math.Abs(float64(target - threeSum))
			if diff < minDiff {
				closest = threeSum
				minDiff = diff
			}
			if target == threeSum {
				break  // 0最接近
			}
			if target > threeSum {
				j++
			} else {
				k--
			}
		}
	}
	return closest
}


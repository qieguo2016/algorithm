/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (46.20%)
 * Likes:    565
 * Dislikes: 0
 * Total Accepted:    25.6K
 * Total Submissions: 55.5K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */

// 双指针解法，与另外一道接雨水类似
// 先判断左右两侧大小，从低的一侧开始走，如果发现有更低的，那么必然可以装起来，因为另一侧有更高的边界
func trap(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1
	leftMax := 0
	rightMax := 0
	for i < j {
		if height[i] < height[j] {
			if height[i] >= leftMax {
				leftMax = height[i]
			} else {
				res += leftMax - height[i]
			}
			i++
		} else {
			if height[j] >= rightMax {
				rightMax = height[j]
			} else {
				res += rightMax - height[j]
			}
			j--
		}
	}
	return res
}


/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (61.94%)
 * Likes:    1195
 * Dislikes: 0
 * Total Accepted:    155K
 * Total Submissions: 250.1K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 * 
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 * 
 * 
 * 
 * 
 * 
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 
 */

// @lc code=start
// 双指针处理，取矮的一侧计算面积，且判断最大面积之后移动矮的一侧指针往前
// 因为面积是由矮侧和宽度决定的，双指针可以保证宽度最大，所以矮侧与其他的搭配都不可能是最大面积
func maxArea(height []int) int {
	// [1,8,6,2,5,4,8,3,7]
	max := 0
	li := 0
	ri := len(height) - 1
	for {
		if ri <= li {
			break
		}
		lc := height[li]
		rc := height[ri]
		w := ri - li
		h := 0
		if rc < lc {
			h = rc
			ri--
		} else {
			h = lc
			li++
		}
		area := w * h
		// fmt.Println(area, w, h)
		if area > max {
			max = area
		}

	}
	return max
}
// @lc code=end


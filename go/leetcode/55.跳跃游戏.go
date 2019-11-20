/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (36.07%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    34K
 * Total Submissions: 93.7K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 判断你是否能够到达最后一个位置。
 * 
 * 示例 1:
 * 
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 * 
 * 
 */

// 与45求最小步数类似，这里不需要关注left
// 只需要关注当前i>=right，当i==right时，那就是说靠跳跃已经跳不过去i这个位置了
func canJump(nums []int) bool {
	right := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] + i > right {
			right = nums[i] + i
		}
		if right >= len(nums) - 1 {
			return true
		}
		if i >= right {
			return false 
		}
	}
	return false
}


/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (31.41%)
 * Likes:    258
 * Dislikes: 0
 * Total Accepted:    17.9K
 * Total Submissions: 56.9K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 * 
 * 示例:
 * 
 * 输入: [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 * 
 * 
 * 说明:
 * 
 * 假设你总是可以到达数组的最后一个位置。
 * 
 */

// 滑动窗口，在每一跳的可选范围里面选包括下一步能走的最远距离
// [5,1,7,9,3,4,1,2,1,1,1,3]
// 第一步在[1,5]之间选，尽量在可选区间的右侧，且下一跳范围更大，也就是i+nums[i]最大
// 第二步更新左右边界
func jump(nums []int) int {
    step, n := 0, len(nums) - 1
    left, right := 0, 0
    for right < n {
        max := right // 尽可能选右边
        for i := left; i <= right; i++ {  // 注意=号，保证可以启动
            if i+nums[i] > max {
                max = i+nums[i]
            }
        }
        left = right+1 // right看过了，下一个
        right = max
        step++
    }
    return step
}


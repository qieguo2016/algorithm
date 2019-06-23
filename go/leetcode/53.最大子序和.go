/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

import (
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := math.MinInt32
	sum := 0
	for _, n := range nums {
		sum += n
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}


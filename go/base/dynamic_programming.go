package base

/**
 * 求最大和子串，求给定整数数组中和最大的连续子串，输出最大和
 * arr: 输入整数数组，至少有一个元素
 */
func GetMaxSum(arr []int) int {
	tmp := 0
	maxSum := 0
	for i := 0; i < len(arr); i++ {
		tmp += arr[i]
		if tmp < 0 {
			tmp = 0
		} else if tmp > maxSum {
			maxSum = tmp
		}
	}
	return maxSum
}

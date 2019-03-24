package base

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

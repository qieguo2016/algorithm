package dp

import (
	"sort"

	"github.com/qieguo2016/algorithm/go/basic/util"
)

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

// 给定金额用最少硬币数兑换
// 求极值一般会考虑使用动态规划。用dp[i]表示i元的最少硬币兑换数，则状态转移方程：
// dp[i]=min(dp[i-c])+1, for c in coins（遍历硬币组合）
// dp[11] = min(dp[6], d[9], dp[10])+1
// dp[1] = min(dp[-4], dp[-1], dp[0]) + 1
// dp[0] = 0
// dp[<0] invalid，用amount+1代替（硬币最小面值为1，最小有amount个，+1标识非法）
// dp有两种模式，一种在迭代的时候增加记忆数组，另外一种是上来先算记忆数组
// PS: 这个兑换有个特例，就是各种硬币是倍数关系的时候，这个时候退化成贪婪模式

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[0]为边界，需要amount+1个
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1 // 默认都非法
	}
	sort.Ints(coins)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				break
			}
			dp[i] = util.Min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

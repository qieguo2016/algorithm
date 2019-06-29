/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

func swap(nums []int, i int, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

func reverse(nums []int, i int, j int)  {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func nextPermutation(nums []int)  {
	// 1 2 4 7 5 3 1 
	// 1 2 5 7 4 3 1
	// 1 2 5 1 3 4 7
	length := len(nums)
	if length < 2 {
		return 
	}
	var i, j int
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j = length-1; j > i; j-- {
				if nums[j] > nums[i] {
					break
				}
			}
			swap(nums, i, j)
			reverse(nums, i+1, length-1)
			break
		}
	}
	if i < 0 {
		reverse(nums, 0, length-1)
	}
}


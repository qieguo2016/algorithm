/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
func plusOne(digits []int) []int {
	n := 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i] + n
		digits[i] = s % 10
		n = s / 10
	}
	ret := []int{}
	if n > 0 {
		ret = append(ret, n)
	}
	ret = append(ret, digits...)
	return ret
}


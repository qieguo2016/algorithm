/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (33.35%)
 * Total Accepted:    35.3K
 * Total Submissions: 105.6K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func leftSub(arr []int, left int) []int {
	if left > len(arr)-1 {
		return []int{}
	}
	return arr[left:]
}

// 在两个有序数组中取第k大的数字
// 将两个数组分段，用排除法逐渐排除掉不可能的分段，分段方式可以按照数组二分，也可以按照k值二分
// 将数组居中2分的时候，要判断m1+n1与k的关系以及nums1[m1]和nums2[n1]的大小关系，从而排除掉大端或者小端
func findKth(arr1 []int, arr2 []int, k int) int {
	// 递归终止条件：某个数组长度为0
	l1, l2 := len(arr1), len(arr2)
	if l1 <= 0 {
		return arr2[k]
	}
	if l2 <= 0 {
		return arr1[k]
	}

	// 注意序号是从0开始的，和个数有-1的差异
	mid1, mid2 := (l1-1)/2, (l2-1)/2

	// 当m/2+n/2大于k的时候，k不可能在最大端, 反之k不可能在最小端
	// :操作是左闭右开，排除大端时为避免将真正的mid元素舍弃掉，所以+1
	if mid1+mid2+1 > k {
		if arr1[mid1] < arr2[mid2] {
			arr2 = arr2[:mid2]
		} else {
			arr1 = arr1[:mid1]
		}
	} else {
		if arr1[mid1] < arr2[mid2] {
			// k - (mid1+1) 与截取操作一致
			k = k - mid1 - 1
			if l1 > 1 {
				arr1 = arr1[mid1+1:]
			} else {
				arr1 = []int{}
			}
		} else {
			k = k - mid2 - 1
			if l2 > 1 {
				arr2 = arr2[mid2+1:]
			} else {
				arr2 = []int{}
			}
		}
	}
	return findKth(arr1, arr2, k)
}

// 将k二分也是同理，这时候要判断k/2是否存在数组中，以及nums1[k/2]和nums2[k/2]的大小关系
func findNth(arr1 []int, arr2 []int, n int) int {
	// 递归终止条件：某个数组长度为0
	l1, l2 := len(arr1), len(arr2)
	if l1 <= 0 {
		return arr2[n]
	}
	if l2 <= 0 {
		return arr1[n]
	}
	if n == 0 {
		return min(arr1[0], arr2[0])
	}
	m1 := (n - 1) / 2
	m2 := (n - 1) / 2
	// n/2就超过了某一个队列的长度，由于n是合法输入，那也不能在长队列的小端
	// 反推一下，假如k在长队列的小端，那么也要从短队列中搬运超过n/2的元素过来，但是n/2大于队列长度，不成立
	if m1 > l1-1 {
		arr2 = leftSub(arr2, m2+1)
		n = n - m2 - 1
	} else if m2 > l2-1 {
		arr1 = leftSub(arr1, m1+1)
		n = n - m1 - 1
		// 下面讨论k/2都在数组中的情况，这时肯定不可能在最小端
	} else if arr1[m1] < arr2[m2] {
		arr1 = leftSub(arr1, m1+1)
		n = n - m1 - 1
	} else {
		arr2 = leftSub(arr2, m2+1)
		n = n - m2 - 1
	}
	return findNth(arr1, arr2, n)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if (l1+l2)%2 == 1 {
		return float64(findNth(nums1, nums2, (l1+l2)/2))
	}
	left := (l1 + l2 - 1) / 2
	right := (l1 + l2) / 2
	return float64(findNth(nums1, nums2, left)+findNth(nums1, nums2, right)) / 2.0
}

package main

import (
	"fmt"
	// "math"
)

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
	if mid1+mid2+1 > k {
		if arr1[mid1] < arr2[mid2] {
			arr2 = arr2[:mid2]
		} else {
			arr1 = arr1[:mid1]
		}
	} else {
		if arr1[mid1] < arr2[mid2] {
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
	if m1 > l1-1 {
		arr2 = leftSub(arr2, m2+1)
		n = n - m2 - 1
	} else if m2 > l2-1 {
		arr1 = leftSub(arr1, m1+1)
		n = n - m1 - 1
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
	left := (l1 + l2 - 1) / 2
	right := (l1 + l2) / 2
	return float64(findKth(nums1, nums2, left)+findKth(nums1, nums2, right)) / 2.0
}

func main() {
	fmt.Println("===== start =====")
	a := []int{1, 5, 7}
	b := []int{2, 4, 6}
	// fmt.Println(a[:1])
	// fmt.Println(b[1:])
	fmt.Println(findNth(a, b, 1))
	fmt.Println(findNth(a, b, 3))
	fmt.Println(findMedianSortedArrays(a, b))
}

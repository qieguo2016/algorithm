package util

import (
	"math/rand"
)

func Swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func MakeRandomArray(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Intn(n * 10)
	}
	return ret
}

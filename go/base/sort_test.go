package base

import (
	"fmt"
	"math/rand"
	"testing"
)

func makeRandomArray(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Intn(n * 10)
	}
	return ret
}

func isArraySort(target []int) bool {
	if len(target) <= 1 {
		return true
	}
	for i := 0; i < len(target)-1; i++ {
		if target[i] > target[i+1] {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := makeRandomArray(10)
		fmt.Printf("origin arr=%v\n", arr)
		ret := QuickSort(arr)
		isSort := isArraySort(ret)
		if !isSort {
			t.Errorf("sort fail, ret=%v\n", ret)
			return
		}
		fmt.Printf("sort success, ret=%v\n", ret)
	}
}

func TestQuickSortInPlace(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := makeRandomArray(10)
		fmt.Printf("origin arr=%v\n", arr)
		ret := QuickSortInPlace(arr)
		isSort := isArraySort(ret)
		if !isSort {
			t.Errorf("sort fail, ret=%v\n", ret)
			return
		}
		fmt.Printf("sort success, ret=%v\n", ret)
	}
}

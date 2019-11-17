package main

import (
	"fmt"
	_ "net/http/pprof"
	// "math"
	// "strconv"
	// "strings"
)


func main() {
	fmt.Println("===== start =====")
	// a := []int{1, 5, 7}
	// b := []int{2, 4, 6}
	// // fmt.Println(a[:1])
	// // fmt.Println(b[1:])
	// fmt.Println(findNth(a, b, 1))
	// fmt.Println(findNth(a, b, 3))
	// fmt.Println(findMedianSortedArrays(a, b))
	// s := "abaab"
	// fmt.Println(s[1:5])
	// fmt.Println(s[1])
	// fmt.Println(longestPalindrome(s))
	// fmt.Println(BubbleSort([]int{1, 9, 4, 3, 8}))
	// fmt.Println(SelectSort([]int{1, 9, 4, 3, 8}))

	// obj := Constructor(1)
	// obj.Put(2, 22)
	// fmt.Println(obj.Get(2)) // 返回  22
	// obj.Put(3, 33)          // 该操作会使得密钥 2 作废
	// fmt.Println(obj.Get(2)) // 返回 -1 (未找到)
	// fmt.Println(obj.Get(3)) // 返回 33

	// n := runtime.NumCPU()
	// fmt.Println("cpu num=", n)
	// runtime.GOMAXPROCS(n)

	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:10000", nil))
	// }()

	// fmt.Println("stage 0, go num=", runtime.NumGoroutine()) // 默认两个go

	// // AlternateOutputViaChannel()
	// // AlternateOutputViaAtomic()
	// base.AlternateOutputViaCond()

	// time.Sleep(100 * time.Second)

	// fmt.Println(helper([]int{1, 2, 3}))
	// fmt.Println(myAtoi("42"))
	// fmt.Println(multiply("456", "123"))
	// strings.IndexOf()
	a := [][]int{
		[]int{5, 1, 9,11},
		[]int{2, 4, 8,10},
		[]int{13, 3, 6, 7},
		[]int{15,14,12,16},
	}
	rotate(a)
	fmt.Println(a)

	fmt.Println("===== end =====")
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ { // 斜线，奇数不转中心，所以小于不等于
		for j := i; j < n-i-1; j++ { // 旋转圈上的一边  1-4 6-1
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n - 1 - j][i]
			matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j]
			matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i]
			matrix[j][n - 1 - i] = tmp
		}
	}
}

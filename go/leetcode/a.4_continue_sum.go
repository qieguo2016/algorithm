package main

// 连续子串和
func subArraySum(arr []int, target int) [][]int {
	s := &subArraySumSol{res: [][]int{}}
	s.call(arr, target)
	return s.res
}

type subArraySumSol struct {
	res [][]int
}

func (s *subArraySumSol) call(arr []int, target int) {
	for i := 0; i < len(arr); i++ {
		s.subArray(arr, target, i)
	}
}

func (s *subArraySumSol) subArray(arr []int, target int, left int) {
	// 找连续子串
	path := []int{}
	for i := left; i < len(arr); i++ {
		target -= arr[i]
		path = append(path, arr[i])
		if target < 0 { // 泛化成连续子串和，如果还有负数则去掉该剪枝
			return
		}
		if target == 0 {
			s.res = append(s.res, append([]int{}, path...))
		}
	}
}

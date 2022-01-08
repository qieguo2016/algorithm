package main

// 找出平均年龄最大的团队
type employerNode struct {
	Val int
	Sub []*employerNode
}

type employerSol struct {
	sum    int
	cnt    int
	oldest float32
}

func (e *employerSol) call(root *employerNode) (sum int, cnt int) {
	if root == nil {
		return 0, 0
	}
	sum, cnt = root.Val, 1
	if len(root.Sub) <= 0 {
		return
	}
	for i := 0; i < len(root.Sub); i++ {
		s, c := e.call(root.Sub[i])
		sum += s
		cnt += c
	}
	if cnt > 1 {
		avg := float32(sum) / float32(cnt)
		if avg > e.oldest {
			e.oldest = avg
		}
	}
	return
}
